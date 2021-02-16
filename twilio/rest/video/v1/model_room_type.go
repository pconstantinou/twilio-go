/*
 * Twilio - Video
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// RoomType the model 'RoomType'
type RoomType string

// List of room_type
const (
	ROOMTYPE_GO RoomType = "go"
	ROOMTYPE_PEER_TO_PEER RoomType = "peer-to-peer"
	ROOMTYPE_GROUP RoomType = "group"
	ROOMTYPE_GROUP_SMALL RoomType = "group-small"
)