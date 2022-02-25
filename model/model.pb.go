// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: model/model.proto

package model

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GitlabProjectModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DefaultBranch   string `protobuf:"bytes,2,opt,name=default_branch,json=defaultBranch,proto3" json:"default_branch,omitempty"`
	Name            string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	GitlabProjectId int64  `protobuf:"varint,4,opt,name=gitlab_project_id,json=gitlabProjectId,proto3" json:"gitlab_project_id,omitempty"`
	Enabled         bool   `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled,omitempty"`
	GlobalEnabled   bool   `protobuf:"varint,6,opt,name=global_enabled,json=globalEnabled,proto3" json:"global_enabled,omitempty"`
	GlobalConfig    string `protobuf:"bytes,7,opt,name=global_config,json=globalConfig,proto3" json:"global_config,omitempty"`
	CreatedAt       string `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       string `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *GitlabProjectModel) Reset() {
	*x = GitlabProjectModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GitlabProjectModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitlabProjectModel) ProtoMessage() {}

func (x *GitlabProjectModel) ProtoReflect() protoreflect.Message {
	mi := &file_model_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitlabProjectModel.ProtoReflect.Descriptor instead.
func (*GitlabProjectModel) Descriptor() ([]byte, []int) {
	return file_model_model_proto_rawDescGZIP(), []int{0}
}

func (x *GitlabProjectModel) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GitlabProjectModel) GetDefaultBranch() string {
	if x != nil {
		return x.DefaultBranch
	}
	return ""
}

func (x *GitlabProjectModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GitlabProjectModel) GetGitlabProjectId() int64 {
	if x != nil {
		return x.GitlabProjectId
	}
	return 0
}

func (x *GitlabProjectModel) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *GitlabProjectModel) GetGlobalEnabled() bool {
	if x != nil {
		return x.GlobalEnabled
	}
	return false
}

func (x *GitlabProjectModel) GetGlobalConfig() string {
	if x != nil {
		return x.GlobalConfig
	}
	return ""
}

func (x *GitlabProjectModel) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *GitlabProjectModel) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type NamespaceModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int64           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name             string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ImagePullSecrets []string        `protobuf:"bytes,3,rep,name=image_pull_secrets,json=imagePullSecrets,proto3" json:"image_pull_secrets,omitempty"`
	CreatedAt        string          `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt        string          `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Projects         []*ProjectModel `protobuf:"bytes,7,rep,name=projects,proto3" json:"projects,omitempty"`
}

func (x *NamespaceModel) Reset() {
	*x = NamespaceModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespaceModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceModel) ProtoMessage() {}

func (x *NamespaceModel) ProtoReflect() protoreflect.Message {
	mi := &file_model_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceModel.ProtoReflect.Descriptor instead.
func (*NamespaceModel) Descriptor() ([]byte, []int) {
	return file_model_model_proto_rawDescGZIP(), []int{1}
}

func (x *NamespaceModel) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NamespaceModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NamespaceModel) GetImagePullSecrets() []string {
	if x != nil {
		return x.ImagePullSecrets
	}
	return nil
}

func (x *NamespaceModel) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *NamespaceModel) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *NamespaceModel) GetProjects() []*ProjectModel {
	if x != nil {
		return x.Projects
	}
	return nil
}

type ProjectModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	GitlabProjectId int64           `protobuf:"varint,3,opt,name=gitlab_project_id,json=gitlabProjectId,proto3" json:"gitlab_project_id,omitempty"`
	GitlabBranch    string          `protobuf:"bytes,4,opt,name=gitlab_branch,json=gitlabBranch,proto3" json:"gitlab_branch,omitempty"`
	GitlabCommit    string          `protobuf:"bytes,5,opt,name=gitlab_commit,json=gitlabCommit,proto3" json:"gitlab_commit,omitempty"`
	Config          string          `protobuf:"bytes,6,opt,name=config,proto3" json:"config,omitempty"`
	OverrideValues  string          `protobuf:"bytes,7,opt,name=override_values,json=overrideValues,proto3" json:"override_values,omitempty"`
	DockerImage     string          `protobuf:"bytes,8,opt,name=docker_image,json=dockerImage,proto3" json:"docker_image,omitempty"`
	PodSelectors    string          `protobuf:"bytes,9,opt,name=pod_selectors,json=podSelectors,proto3" json:"pod_selectors,omitempty"`
	NamespaceId     int64           `protobuf:"varint,10,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
	Atomic          bool            `protobuf:"varint,11,opt,name=atomic,proto3" json:"atomic,omitempty"`
	CreatedAt       string          `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       string          `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	ExtraValues     string          `protobuf:"bytes,14,opt,name=extra_values,json=extraValues,proto3" json:"extra_values,omitempty"`
	Namespace       *NamespaceModel `protobuf:"bytes,15,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *ProjectModel) Reset() {
	*x = ProjectModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectModel) ProtoMessage() {}

func (x *ProjectModel) ProtoReflect() protoreflect.Message {
	mi := &file_model_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectModel.ProtoReflect.Descriptor instead.
func (*ProjectModel) Descriptor() ([]byte, []int) {
	return file_model_model_proto_rawDescGZIP(), []int{2}
}

func (x *ProjectModel) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProjectModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProjectModel) GetGitlabProjectId() int64 {
	if x != nil {
		return x.GitlabProjectId
	}
	return 0
}

func (x *ProjectModel) GetGitlabBranch() string {
	if x != nil {
		return x.GitlabBranch
	}
	return ""
}

func (x *ProjectModel) GetGitlabCommit() string {
	if x != nil {
		return x.GitlabCommit
	}
	return ""
}

func (x *ProjectModel) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *ProjectModel) GetOverrideValues() string {
	if x != nil {
		return x.OverrideValues
	}
	return ""
}

func (x *ProjectModel) GetDockerImage() string {
	if x != nil {
		return x.DockerImage
	}
	return ""
}

func (x *ProjectModel) GetPodSelectors() string {
	if x != nil {
		return x.PodSelectors
	}
	return ""
}

func (x *ProjectModel) GetNamespaceId() int64 {
	if x != nil {
		return x.NamespaceId
	}
	return 0
}

func (x *ProjectModel) GetAtomic() bool {
	if x != nil {
		return x.Atomic
	}
	return false
}

func (x *ProjectModel) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ProjectModel) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *ProjectModel) GetExtraValues() string {
	if x != nil {
		return x.ExtraValues
	}
	return ""
}

func (x *ProjectModel) GetNamespace() *NamespaceModel {
	if x != nil {
		return x.Namespace
	}
	return nil
}

type FileModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Path          string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Size          uint64 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Username      string `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Namespace     string `protobuf:"bytes,5,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Pod           string `protobuf:"bytes,6,opt,name=pod,proto3" json:"pod,omitempty"`
	Container     string `protobuf:"bytes,7,opt,name=container,proto3" json:"container,omitempty"`
	ContainerPath string `protobuf:"bytes,8,opt,name=container_path,json=containerPath,proto3" json:"container_path,omitempty"`
	CreatedAt     string `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt     string `protobuf:"bytes,11,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	IsDeleted     bool   `protobuf:"varint,12,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
}

func (x *FileModel) Reset() {
	*x = FileModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileModel) ProtoMessage() {}

func (x *FileModel) ProtoReflect() protoreflect.Message {
	mi := &file_model_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileModel.ProtoReflect.Descriptor instead.
func (*FileModel) Descriptor() ([]byte, []int) {
	return file_model_model_proto_rawDescGZIP(), []int{3}
}

func (x *FileModel) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FileModel) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *FileModel) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *FileModel) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *FileModel) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *FileModel) GetPod() string {
	if x != nil {
		return x.Pod
	}
	return ""
}

func (x *FileModel) GetContainer() string {
	if x != nil {
		return x.Container
	}
	return ""
}

func (x *FileModel) GetContainerPath() string {
	if x != nil {
		return x.ContainerPath
	}
	return ""
}

func (x *FileModel) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *FileModel) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *FileModel) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

func (x *FileModel) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

var File_model_model_proto protoreflect.FileDescriptor

var file_model_model_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x02, 0x0a, 0x12, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x5f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x67, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xcb, 0x01, 0x0a, 0x0e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x50,
	0x75, 0x6c, 0x6c, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x29, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x22, 0xfc, 0x03, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x67, 0x69, 0x74, 0x6c,
	0x61, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x5f, 0x62,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x67, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x23, 0x0a, 0x0d, 0x67, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69,
	0x64, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x6f, 0x64, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x6f, 0x64, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x74,
	0x6f, 0x6d, 0x69, 0x63, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x74, 0x6f, 0x6d,
	0x69, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x74, 0x72, 0x61, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x22, 0xd0, 0x02, 0x0a, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x70, 0x6f, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x75, 0x63, 0x2d, 0x63, 0x6e, 0x7a, 0x6a, 0x2f, 0x6d, 0x61, 0x72,
	0x73, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x33, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x3b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_model_proto_rawDescOnce sync.Once
	file_model_model_proto_rawDescData = file_model_model_proto_rawDesc
)

func file_model_model_proto_rawDescGZIP() []byte {
	file_model_model_proto_rawDescOnce.Do(func() {
		file_model_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_model_proto_rawDescData)
	})
	return file_model_model_proto_rawDescData
}

var file_model_model_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_model_model_proto_goTypes = []interface{}{
	(*GitlabProjectModel)(nil), // 0: GitlabProjectModel
	(*NamespaceModel)(nil),     // 1: NamespaceModel
	(*ProjectModel)(nil),       // 2: ProjectModel
	(*FileModel)(nil),          // 3: FileModel
}
var file_model_model_proto_depIdxs = []int32{
	2, // 0: NamespaceModel.projects:type_name -> ProjectModel
	1, // 1: ProjectModel.namespace:type_name -> NamespaceModel
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_model_model_proto_init() }
func file_model_model_proto_init() {
	if File_model_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GitlabProjectModel); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespaceModel); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectModel); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileModel); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_model_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_model_proto_goTypes,
		DependencyIndexes: file_model_model_proto_depIdxs,
		MessageInfos:      file_model_model_proto_msgTypes,
	}.Build()
	File_model_model_proto = out.File
	file_model_model_proto_rawDesc = nil
	file_model_model_proto_goTypes = nil
	file_model_model_proto_depIdxs = nil
}
