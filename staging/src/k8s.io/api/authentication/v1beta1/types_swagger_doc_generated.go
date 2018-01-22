/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_TokenReview = map[string]string{
	"":       "TokenReview attempts to authenticate a token to a known user. Note: TokenReview requests may be cached by the webhook token authenticator plugin in the kube-apiserver.",
	"spec":   "Spec holds information about the request being evaluated",
	"status": "Status is filled in by the server and indicates whether the request can be authenticated.",
}

func (TokenReview) SwaggerDoc() map[string]string {
	return map_TokenReview
}

var map_TokenReviewSpec = map[string]string{
	"":      "TokenReviewSpec is a description of the token authentication request.",
	"token": "Token is the opaque bearer token.",
}

func (TokenReviewSpec) SwaggerDoc() map[string]string {
	return map_TokenReviewSpec
}

var map_TokenReviewStatus = map[string]string{
	"":              "TokenReviewStatus is the result of the token authentication request.",
	"authenticated": "Authenticated indicates that the token was associated with a known user.",
	"user":          "User is the UserInfo associated with the provided token.",
	"error":         "Error indicates that the token couldn't be checked",
}

func (TokenReviewStatus) SwaggerDoc() map[string]string {
	return map_TokenReviewStatus
}

var map_UserInfo = map[string]string{
	"":         "UserInfo holds the information about the user needed to implement the user.Info interface.",
	"username": "The name that uniquely identifies this user among all active users.",
	"uid":      "A unique value that identifies this user across time. If this user is deleted and another user by the same name is added, they will have different UIDs.",
	"groups":   "The names of groups this user is a part of.",
	"extra":    "Any additional information provided by the authenticator.",
}

func (UserInfo) SwaggerDoc() map[string]string {
	return map_UserInfo
}

// AUTO-GENERATED FUNCTIONS END HERE
