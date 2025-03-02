// Copyright 2022 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file https://github.com/redpanda-data/redpanda/blob/dev/licenses/bsl.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package api

import (
	_ "context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	_ "time"

	"github.com/twmb/franz-go/pkg/kmsg"
	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"

	"github.com/cloudhut/common/rest"
	"github.com/cloudhut/kowl/backend/pkg/console"
	"github.com/go-chi/chi"
)

func (api *API) handleGetTopics() http.HandlerFunc {
	type response struct {
		Topics []*console.TopicSummary `json:"topics"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		topics, err := api.ConsoleSvc.GetTopicsOverview(r.Context())
		if err != nil {
			restErr := &rest.Error{
				Err:      err,
				Status:   http.StatusInternalServerError,
				Message:  "Could not list topics from Kafka cluster",
				IsSilent: false,
			}
			rest.SendRESTError(w, r, api.Logger, restErr)
			return
		}

		visibleTopics := make([]*console.TopicSummary, 0, len(topics))
		for _, topic := range topics {
			// Check if logged in user is allowed to see this topic. If not remove the topic from the list.
			canSee, restErr := api.Hooks.Console.CanSeeTopic(r.Context(), topic.TopicName)
			if restErr != nil {
				rest.SendRESTError(w, r, api.Logger, restErr)
				return
			}

			if canSee {
				visibleTopics = append(visibleTopics, topic)
			}

			// Attach allowed actions for each topic
			topic.AllowedActions, restErr = api.Hooks.Console.AllowedTopicActions(r.Context(), topic.TopicName)
			if restErr != nil {
				rest.SendRESTError(w, r, api.Logger, restErr)
				return
			}
		}

		response := response{
			Topics: visibleTopics,
		}
		rest.SendResponse(w, r, api.Logger, http.StatusOK, response)
	}
}

// handleGetPartitions returns an overview of all partitions and their watermarks in the given topic
func (api *API) handleGetPartitions() http.HandlerFunc {
	type response struct {
		TopicName  string                          `json:"topicName"`
		Partitions []console.TopicPartitionDetails `json:"partitions"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		topicName := chi.URLParam(r, "topicName")
		logger := api.Logger.With(zap.String("topic_name", topicName))

		// Check if logged in user is allowed to view partitions for the given topic
		canView, restErr := api.Hooks.Console.CanViewTopicPartitions(r.Context(), topicName)
		if restErr != nil {
			rest.SendRESTError(w, r, logger, restErr)
			return
		}
		if !canView {
			restErr := &rest.Error{
				Err:      fmt.Errorf("requester has no permissions to view partitions for the requested topic"),
				Status:   http.StatusForbidden,
				Message:  "You don't have permissions to view partitions for that topic",
				IsSilent: false,
			}
			rest.SendRESTError(w, r, logger, restErr)
			return
		}

		topicDetails, restErr := api.ConsoleSvc.GetTopicDetails(r.Context(), []string{topicName})
		if restErr != nil {
			rest.SendRESTError(w, r, logger, restErr)
			return
		}

		if len(topicDetails) != 1 {
			restErr := &rest.Error{
				Err:      fmt.Errorf("expected exactly one topic detail in response, but got '%d'", len(topicDetails)),
				Status:   http.StatusInternalServerError,
				Message:  "Internal server error in Kowl, please file a issue in GitHub if you face this issue. The backend logs will contain more information.",
				IsSilent: false,
			}
			rest.SendRESTError(w, r, logger, restErr)
			return
		}

		res := response{
			TopicName:  topicName,
			Partitions: topicDetails[0].Partitions,
		}
		rest.SendResponse(w, r, logger, http.StatusOK, res)
	}
}

// handleGetTopicConfig returns all set configuration options for a specific topic
func (api *API) handleGetTopicConfig() http.HandlerFunc {
	type response struct {
		TopicDescription *console.TopicConfig `json:"topicDescription"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		topicName := chi.URLParam(r, "topicName")
		logger := api.Logger.With(zap.String("topic_name", topicName))

		// Check if logged in user is allowed to view partitions for the given topic
		canView, restErr := api.Hooks.Console.CanViewTopicConfig(r.Context(), topicName)
		if restErr != nil {
			rest.SendRESTError(w, r, logger, restErr)
			return
		}
		if !canView {
			restErr := &rest.Error{
				Err:      fmt.Errorf("requester has no permissions to view config for the requested topic"),
				Status:   http.StatusForbidden,
				Message:  "You don't have permissions to view the config for that topic",
				IsSilent: false,
			}
			rest.SendRESTError(w, r, logger, restErr)
			return
		}

		description, restErr := api.ConsoleSvc.GetTopicConfigs(r.Context(), topicName, nil)
		if restErr != nil {
			rest.SendRESTError(w, r, logger, restErr)
			return
		}

		res := response{
			TopicDescription: description,
		}
		rest.SendResponse(w, r, api.Logger, http.StatusOK, res)
	}
}

func (api *API) handleDeleteTopic() http.HandlerFunc {
	type response struct {
		Status string `json:"status"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		topicName := chi.URLParam(r, "topicName")

		// Check if logged in user is allowed to view partitions for the given topic
		canDelete, restErr := api.Hooks.Console.CanDeleteTopic(r.Context(), topicName)
		if restErr != nil {
			rest.SendRESTError(w, r, api.Logger, restErr)
			return
		}
		if !canDelete {
			restErr := &rest.Error{
				Err:      fmt.Errorf("requester has no permissions to delete this topic"),
				Status:   http.StatusForbidden,
				Message:  "You don't have permissions to delete this topic",
				IsSilent: false,
			}
			rest.SendRESTError(w, r, api.Logger, restErr)
			return
		}

		restErr = api.ConsoleSvc.DeleteTopic(r.Context(), topicName)
		if restErr != nil {
			rest.SendRESTError(w, r, api.Logger, restErr)
			return
		}

		rest.SendResponse(w, r, api.Logger, http.StatusOK, response{Status: "Success"})
	}
}

type deleteTopicRecordsRequest struct {
	// Partitions contains partitions to delete records from.
	Partitions []struct {
		// Partition is a partition to delete records from.
		Partition int32 `json:"partitionId"`

		// Offset is the offset to set the partition's low watermark (start
		// offset) to. After a successful response, all records before this
		// offset are considered deleted and are no longer readable.
		//
		// To delete all records, use -1, which is mapped to the partition's
		// current high watermark.
		Offset int64 `json:"offset"`
	} `json:"partitions"`
}

func (d *deleteTopicRecordsRequest) OK() error {
	if len(d.Partitions) == 0 {
		return fmt.Errorf("at least one partition must be specified")
	}

	for _, partition := range d.Partitions {
		if partition.Offset < -1 {
			return fmt.Errorf("partition offset must be greater than -1")
		}
	}

	return nil
}

func (api *API) handleDeleteTopicRecords() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		topicName := chi.URLParam(r, "topicName")

		// 1. Parse and validate request
		var req deleteTopicRecordsRequest
		restErr := rest.Decode(w, r, &req)
		if restErr != nil {
			rest.SendRESTError(w, r, api.Logger, restErr)
			return
		}

		// 2. Check if logged in user is allowed to view partitions for the given topic
		canDelete, restErr := api.Hooks.Console.CanDeleteTopicRecords(r.Context(), topicName)
		if restErr != nil {
			rest.SendRESTError(w, r, api.Logger, restErr)
			return
		}
		if !canDelete {
			restErr := &rest.Error{
				Err:      fmt.Errorf("requester has no permissions to delete records in this topic"),
				Status:   http.StatusForbidden,
				Message:  "You don't have permissions to delete this topic",
				IsSilent: false,
			}
			rest.SendRESTError(w, r, api.Logger, restErr)
			return
		}

		// 3. Submit delete topic records request
		deleteReq := kmsg.NewDeleteRecordsRequestTopic()
		deleteReq.Topic = topicName
		deleteReq.Partitions = make([]kmsg.DeleteRecordsRequestTopicPartition, len(req.Partitions))
		for i, partition := range req.Partitions {
			pReq := kmsg.NewDeleteRecordsRequestTopicPartition()
			pReq.Partition = partition.Partition
			pReq.Offset = partition.Offset
			deleteReq.Partitions[i] = pReq
		}

		deleteRes, restErr := api.ConsoleSvc.DeleteTopicRecords(r.Context(), deleteReq)
		if restErr != nil {
			rest.SendRESTError(w, r, api.Logger, restErr)
			return
		}

		rest.SendResponse(w, r, api.Logger, http.StatusOK, deleteRes)
	}
}

// handleGetTopicsConfigs returns all set configuration options for one or more topics
func (api *API) handleGetTopicsConfigs() http.HandlerFunc {
	type response struct {
		TopicDescriptions []*console.TopicConfig `json:"topicDescriptions"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// 1. Parse optional filters. If no filter is set they will be treated as wildcards
		var topicNames []string
		requestedTopicNames := r.URL.Query().Get("topicNames")
		if requestedTopicNames != "" {
			topicNames = strings.Split(requestedTopicNames, ",")
		}

		var configKeys []string
		requestedConfigKeys := r.URL.Query().Get("configKeys")
		if requestedConfigKeys != "" {
			configKeys = strings.Split(requestedConfigKeys, ",")
		}

		logger := api.Logger.With(zap.Int("topics_length", len(topicNames)), zap.Int("config_keys_length", len(configKeys)))

		// 2. Fetch all topic names from metadata as no topic filter has been specified
		if len(topicNames) == 0 {
			var err error
			topicNames, err = api.ConsoleSvc.GetAllTopicNames(r.Context(), nil)
			if err != nil {
				restErr := &rest.Error{
					Err:      fmt.Errorf("failed to request metadata to fetch topic names: %w", err),
					Status:   http.StatusForbidden,
					Message:  fmt.Sprintf("Failed to fetch metadata from brokers to fetch topicNames '%v'", err.Error()),
					IsSilent: false,
				}
				rest.SendRESTError(w, r, logger, restErr)
				return
			}
		}

		// 3. Check if user is allowed to view the config for these topics
		for _, topicName := range topicNames {
			canView, restErr := api.Hooks.Console.CanViewTopicConfig(r.Context(), topicName)
			if restErr != nil {
				rest.SendRESTError(w, r, logger, restErr)
				return
			}
			if !canView {
				restErr := &rest.Error{
					Err:      fmt.Errorf("requester has no permissions to view config for one of the requested topics"),
					Status:   http.StatusForbidden,
					Message:  fmt.Sprintf("You don't have permissions to view the config for topic '%v'", topicName),
					IsSilent: false,
				}
				rest.SendRESTError(w, r, logger, restErr)
				return
			}
		}

		// 4. Request topics configs and return them
		descriptions, err := api.ConsoleSvc.GetTopicsConfigs(r.Context(), topicNames, configKeys)
		if err != nil {
			restErr := &rest.Error{
				Err:      fmt.Errorf("failed to describe topic configs: %w", err),
				Status:   http.StatusServiceUnavailable,
				Message:  fmt.Sprintf("Failed to describe topic configs '%v'", err.Error()),
				IsSilent: false,
			}
			rest.SendRESTError(w, r, logger, restErr)
			return
		}
		result := make([]*console.TopicConfig, 0, len(descriptions))
		for _, description := range descriptions {
			result = append(result, description)
		}

		res := response{
			TopicDescriptions: result,
		}
		rest.SendResponse(w, r, api.Logger, http.StatusOK, res)
	}
}

// handleGetTopicConsumers returns all consumers along with their summed lag which consume the given topic
func (api *API) handleGetTopicConsumers() http.HandlerFunc {
	type response struct {
		TopicName string                        `json:"topicName"`
		Consumers []*console.TopicConsumerGroup `json:"topicConsumers"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		topicName := chi.URLParam(r, "topicName")
		logger := api.Logger.With(zap.String("topic_name", topicName))

		// Check if logged in user is allowed to view partitions for the given topic
		canView, restErr := api.Hooks.Console.CanViewTopicConsumers(r.Context(), topicName)
		if restErr != nil {
			rest.SendRESTError(w, r, logger, restErr)
			return
		}
		if !canView {
			restErr := &rest.Error{
				Err:      fmt.Errorf("requester has no permissions to view topic consumers for the requested topic"),
				Status:   http.StatusForbidden,
				Message:  "You don't have permissions to view the config for that topic",
				IsSilent: false,
			}
			rest.SendRESTError(w, r, logger, restErr)
			return
		}

		consumers, err := api.ConsoleSvc.ListTopicConsumers(r.Context(), topicName)
		if err != nil {
			restErr := &rest.Error{
				Err:      err,
				Status:   http.StatusInternalServerError,
				Message:  "Could not list topic consumers for requested topic",
				IsSilent: false,
			}
			rest.SendRESTError(w, r, logger, restErr)
			return
		}

		res := response{
			TopicName: topicName,
			Consumers: consumers,
		}
		rest.SendResponse(w, r, logger, http.StatusOK, res)
	}
}

func (api *API) handleGetTopicsOffsets() http.HandlerFunc {
	type response struct {
		TopicOffsets []console.TopicOffset `json:"topicOffsets"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// 1. Parse topic names and timestamp from URL. It's a list of topic names that is comma separated
		requestedTopicNames := r.URL.Query().Get("topicNames")
		if requestedTopicNames == "" {
			restErr := &rest.Error{
				Err:      fmt.Errorf("required parameter topicNames is missing"),
				Status:   http.StatusBadRequest,
				Message:  "Required parameter topicNames is missing",
				IsSilent: false,
			}
			rest.SendRESTError(w, r, api.Logger, restErr)
			return
		}
		topicNames := strings.Split(requestedTopicNames, ",")

		timestampStr := r.URL.Query().Get("timestamp")
		if timestampStr == "" {
			restErr := &rest.Error{
				Err:      fmt.Errorf("required parameter timestamp is missing"),
				Status:   http.StatusBadRequest,
				Message:  "Required parameter timestamp is missing",
				IsSilent: false,
			}
			rest.SendRESTError(w, r, api.Logger, restErr)
			return
		}
		timestamp, err := strconv.Atoi(timestampStr)
		if err != nil {
			restErr := &rest.Error{
				Err:      fmt.Errorf("timestamp parameter must be a valid int"),
				Status:   http.StatusBadRequest,
				Message:  "Timestamp parameter must be a valid int",
				IsSilent: false,
			}
			rest.SendRESTError(w, r, api.Logger, restErr)
			return
		}

		// 2. Check if logged in user is allowed list partitions (always true for Kowl, but not for Kowl Business)
		for _, topic := range topicNames {
			canView, restErr := api.Hooks.Console.CanViewTopicPartitions(r.Context(), topic)
			if restErr != nil {
				rest.SendRESTError(w, r, api.Logger, restErr)
				return
			}

			if !canView {
				restErr := &rest.Error{
					Err:          fmt.Errorf("requester has no permissions to view partitions for the requested topic"),
					Status:       http.StatusForbidden,
					Message:      "You don't have permissions to view partitions for that topic",
					IsSilent:     false,
					InternalLogs: []zapcore.Field{zap.String("topic_name", topic)},
				}
				rest.SendRESTError(w, r, api.Logger, restErr)
				return
			}
		}

		// 3. Request topic
		topicOffsets, err := api.ConsoleSvc.ListOffsets(r.Context(), topicNames, int64(timestamp))
		if err != nil {
			restErr := &rest.Error{
				Err:      fmt.Errorf("failed to list offsets: %w", err),
				Status:   http.StatusForbidden,
				Message:  fmt.Sprintf("Failed to list offsets from Kafka: %v", err.Error()),
				IsSilent: false,
			}
			rest.SendRESTError(w, r, api.Logger, restErr)
			return
		}

		res := response{
			TopicOffsets: topicOffsets,
		}
		rest.SendResponse(w, r, api.Logger, http.StatusOK, res)
	}
}
