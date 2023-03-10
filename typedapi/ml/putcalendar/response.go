// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/1ad7fe36297b3a8e187b2259dedaf68a47bc236e

package putcalendar

// Response holds the response body struct for the package putcalendar
//
// https://github.com/elastic/elasticsearch-specification/blob/1ad7fe36297b3a8e187b2259dedaf68a47bc236e/specification/ml/put_calendar/MlPutCalendarResponse.ts#L22-L31

type Response struct {

	// CalendarId A string that uniquely identifies a calendar.
	CalendarId string `json:"calendar_id"`
	// Description A description of the calendar.
	Description *string `json:"description,omitempty"`
	// JobIds A list of anomaly detection job identifiers or group names.
	JobIds []string `json:"job_ids"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}