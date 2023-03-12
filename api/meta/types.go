/*
Copyright 2021 The Pixiu Authors.

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

package meta

type IdMeta struct {
	Id int64 `uri:"id" binding:"required"`
}

type CloudMeta struct {
	Cloud string `uri:"cloud_name" binding:"required"`
}

type KubeMeta struct {
	Namespace  string `uri:"namespace" binding:"required"`
	ObjectName string `uri:"object_name" binding:"required"`
}

type NamespaceMeta struct {
	CloudMeta  `json:",inline"`
	ObjectName string `uri:"object_name" binding:"required"`
}

type CreateOptions struct {
	CloudMeta `json:",inline"`

	Namespace string `uri:"namespace" binding:"required"`
}

type UpdateOptions struct {
	CloudMeta `json:",inline"`

	Namespace  string `uri:"namespace" binding:"required"`
	ObjectName string `uri:"object_name" binding:"required"`
}

type DeleteOptions struct {
	CloudMeta `json:",inline"`

	Namespace  string `uri:"namespace" binding:"required"`
	ObjectName string `uri:"object_name" binding:"required"`
}

type GetOptions struct {
	CloudMeta `json:",inline"`

	Namespace  string `uri:"namespace" binding:"required"`
	ObjectName string `uri:"object_name" binding:"required"`
}

type ListOptions struct {
	CloudMeta `json:",inline"`

	Namespace string `uri:"namespace" binding:"required"`

	Selector *ListSelector
}
