// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20190819

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Acl struct {

	// Acl资源类型，（0:UNKNOWN，1:ANY，2:TOPIC，3:GROUP，4:CLUSTER，5:TRANSACTIONAL_ID）当前只有TOPIC，
	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// 资源名称，和resourceType相关如当resourceType为TOPIC时，则该字段表示topic名称，当resourceType为GROUP时，该字段表示group名称
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 用户列表，默认为User:*，表示任何user都可以访问，当前用户只能是用户列表中包含的用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	Principal *string `json:"Principal,omitempty" name:"Principal"`

	// 默认为*，表示任何host都可以访问，当前ckafka不支持host为*，但是后面开源kafka的产品化会直接支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	Host *string `json:"Host,omitempty" name:"Host"`

	// Acl操作方式(0:UNKNOWN，1:ANY，2:ALL，3:READ，4:WRITE，5:CREATE，6:DELETE，7:ALTER，8:DESCRIBE，9:CLUSTER_ACTION，10:DESCRIBE_CONFIGS，11:ALTER_CONFIGS，12:IDEMPOTEN_WRITE)
	Operation *int64 `json:"Operation,omitempty" name:"Operation"`

	// 权限类型(0:UNKNOWN，1:ANY，2:DENY，3:ALLOW)
	PermissionType *int64 `json:"PermissionType,omitempty" name:"PermissionType"`
}

type AclResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的总数据条数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// ACL列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		AclList []*Acl `json:"AclList,omitempty" name:"AclList" list`
	} `json:"Response"`
}

func (r *AclResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AclResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AppIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合要求的所有AppId数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 符合要求的App Id列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		AppIdList []*int64 `json:"AppIdList,omitempty" name:"AppIdList" list`
	} `json:"Response"`
}

func (r *AppIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AppIdResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Assignment struct {

	// assingment版本信息
	Version *int64 `json:"Version,omitempty" name:"Version"`

	// topic信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Topics []*GroupInfoTopics `json:"Topics,omitempty" name:"Topics" list`
}

type Config struct {

	// 消息保留时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Retention *int64 `json:"Retention,omitempty" name:"Retention"`

	// 最小同步复制数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitempty" name:"MinInsyncReplicas"`

	// 日志清理模式，默认 delete。
	// delete：日志按保存时间删除；compact：日志按 key 压缩；compact, delete：日志按 key 压缩且会保存时间删除。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CleanUpPolicy *string `json:"CleanUpPolicy,omitempty" name:"CleanUpPolicy"`

	// Segment 分片滚动的时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentMs *int64 `json:"SegmentMs,omitempty" name:"SegmentMs"`

	// 0表示 false。 1表示 true。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitempty" name:"UncleanLeaderElectionEnable"`

	// Segment 分片滚动的字节数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentBytes *int64 `json:"SegmentBytes,omitempty" name:"SegmentBytes"`

	// 最大消息字节数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxMessageBytes *int64 `json:"MaxMessageBytes,omitempty" name:"MaxMessageBytes"`
}

type ConsumerGroup struct {

	// 用户组名称
	ConsumerGroupName *string `json:"ConsumerGroupName,omitempty" name:"ConsumerGroupName"`

	// 订阅信息实体
	SubscribedInfo []*SubscribedInfo `json:"SubscribedInfo,omitempty" name:"SubscribedInfo" list`
}

type ConsumerGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的消费组数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 主题列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		TopicList []*ConsumerGroupTopic `json:"TopicList,omitempty" name:"TopicList" list`

		// 消费分组List
	// 注意：此字段可能返回 null，表示取不到有效值。
		GroupList []*ConsumerGroup `json:"GroupList,omitempty" name:"GroupList" list`

		// 所有分区数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalPartition *int64 `json:"TotalPartition,omitempty" name:"TotalPartition"`

		// 监控的分区列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		PartitionListForMonitor []*Partition `json:"PartitionListForMonitor,omitempty" name:"PartitionListForMonitor" list`

		// 主题总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalTopic *int64 `json:"TotalTopic,omitempty" name:"TotalTopic"`

		// 监控的主题列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		TopicListForMonitor []*ConsumerGroupTopic `json:"TopicListForMonitor,omitempty" name:"TopicListForMonitor" list`

		// 监控的组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		GroupListForMonitor []*Group `json:"GroupListForMonitor,omitempty" name:"GroupListForMonitor" list`
	} `json:"Response"`
}

func (r *ConsumerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ConsumerGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ConsumerGroupTopic struct {

	// 主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type CreateAclRequest struct {
	*tchttp.BaseRequest

	// 实例id信息
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Acl资源类型，（0:UNKNOWN，1:ANY，2:TOPIC，3:GROUP，4:CLUSTER，5:TRANSACTIONAL_ID）当前只有TOPIC，其它字段用于后续兼容开源kafka的acl时使用
	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// 资源名称，和resourceType相关如当resourceType为TOPIC时，则该字段表示topic名称，当resourceType为GROUP时，该字段表示group名称
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// Acl操作方式(0:UNKNOWN，1:ANY，2:ALL，3:READ，4:WRITE，5:CREATE，6:DELETE，7:ALTER，8:DESCRIBE，9:CLUSTER_ACTION，10:DESCRIBE_CONFIGS，11:ALTER_CONFIGS，
	Operation *int64 `json:"Operation,omitempty" name:"Operation"`

	// 权限类型(0:UNKNOWN，1:ANY，2:DENY，3:ALLOW)，当前ckakfa支持ALLOW(相当于白名单)，其它用于后续兼容开源kafka的acl时使用
	PermissionType *int64 `json:"PermissionType,omitempty" name:"PermissionType"`

	// 默认为*，表示任何host都可以访问，当前ckafka不支持host为*，但是后面开源kafka的产品化会直接支持
	Host *string `json:"Host,omitempty" name:"Host"`

	// 用户列表，默认为*，表示任何user都可以访问，当前用户只能是用户列表中包含的用户
	Principal *string `json:"Principal,omitempty" name:"Principal"`
}

func (r *CreateAclRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAclRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAclResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果
		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAclResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAclResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePartitionRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 主题分区个数
	PartitionNum *int64 `json:"PartitionNum,omitempty" name:"PartitionNum"`
}

func (r *CreatePartitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePartitionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePartitionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的结果集
		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePartitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePartitionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTopicIpWhiteListRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// ip白名单列表
	IpWhiteList []*string `json:"IpWhiteList,omitempty" name:"IpWhiteList" list`
}

func (r *CreateTopicIpWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTopicIpWhiteListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTopicIpWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除主题IP白名单结果
		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTopicIpWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTopicIpWhiteListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTopicRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主题名称，是一个不超过 64 个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Partition个数，大于0
	PartitionNum *int64 `json:"PartitionNum,omitempty" name:"PartitionNum"`

	// 副本个数，不能多于 broker 数，最大为3
	ReplicaNum *int64 `json:"ReplicaNum,omitempty" name:"ReplicaNum"`

	// ip白名单开关, 1:打开  0:关闭，默认不打开
	EnableWhiteList *int64 `json:"EnableWhiteList,omitempty" name:"EnableWhiteList"`

	// Ip白名单列表，配额限制，enableWhileList=1时必选
	IpWhiteList []*string `json:"IpWhiteList,omitempty" name:"IpWhiteList" list`

	// 清理日志策略，日志清理模式，默认为"delete"。"delete"：日志按保存时间删除，"compact"：日志按 key 压缩，"compact, delete"：日志按 key 压缩且会按保存时间删除。
	CleanUpPolicy *string `json:"CleanUpPolicy,omitempty" name:"CleanUpPolicy"`

	// 主题备注，是一个不超过 64 个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)
	Note *string `json:"Note,omitempty" name:"Note"`

	// 默认为1
	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitempty" name:"MinInsyncReplicas"`

	// 是否允许未同步的副本选为leader，false:不允许，true:允许，默认不允许
	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitempty" name:"UncleanLeaderElectionEnable"`

	// 可消息选。保留时间，单位ms，当前最小值为60000ms
	RetentionMs *int64 `json:"RetentionMs,omitempty" name:"RetentionMs"`

	// Segment分片滚动的时长，单位ms，当前最小为3600000ms
	SegmentMs *int64 `json:"SegmentMs,omitempty" name:"SegmentMs"`
}

func (r *CreateTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTopicRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTopicResp struct {

	// 主题Id
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

type CreateTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回创建结果
		Result *CreateTopicResp `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTopicResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateUserRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用户名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用户密码
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *CreateUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateUserRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的结果
		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateUserResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAclRequest struct {
	*tchttp.BaseRequest

	// 实例id信息
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Acl资源类型，（0:UNKNOWN，1:ANY，2:TOPIC，3:GROUP，4:CLUSTER，5:TRANSACTIONAL_ID）当前只有TOPIC，其它字段用于后续兼容开源kafka的acl时使用
	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// 资源名称，和resourceType相关如当resourceType为TOPIC时，则该字段表示topic名称，当resourceType为GROUP时，该字段表示group名称
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// Acl操作方式(0:UNKNOWN，1:ANY，2:ALL，3:READ，4:WRITE，5:CREATE，6:DELETE，7:ALTER，8:DESCRIBE，9:CLUSTER_ACTION，10:DESCRIBE_CONFIGS，11:ALTER_CONFIGS，12:IDEMPOTEN_WRITE)，当前ckafka只支持READ,WRITE，其它用于后续兼容开源kafka的acl时使用
	Operation *int64 `json:"Operation,omitempty" name:"Operation"`

	// 权限类型(0:UNKNOWN，1:ANY，2:DENY，3:ALLOW)，当前ckakfa支持ALLOW(相当于白名单)，其它用于后续兼容开源kafka的acl时使用
	PermissionType *int64 `json:"PermissionType,omitempty" name:"PermissionType"`

	// 默认为*，表示任何host都可以访问，当前ckafka不支持host为*，但是后面开源kafka的产品化会直接支持
	Host *string `json:"Host,omitempty" name:"Host"`

	// 用户列表，默认为*，表示任何user都可以访问，当前用户只能是用户列表中包含的用户
	Principal *string `json:"Principal,omitempty" name:"Principal"`
}

func (r *DeleteAclRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAclRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAclResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果
		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAclResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAclResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicIpWhiteListRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// ip白名单列表
	IpWhiteList []*string `json:"IpWhiteList,omitempty" name:"IpWhiteList" list`
}

func (r *DeleteTopicIpWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTopicIpWhiteListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicIpWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除主题IP白名单结果
		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTopicIpWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTopicIpWhiteListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicRequest struct {
	*tchttp.BaseRequest

	// ckafka 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// ckafka 主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

func (r *DeleteTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTopicRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的结果集
		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTopicResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteUserRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用户名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteUserRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果
		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteUserResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeACLRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Acl资源类型，（0:UNKNOWN，1:ANY，2:TOPIC，3:GROUP，4:CLUSTER，5:TRANSACTIONAL_ID）当前只有TOPIC，其它字段用于后续兼容开源kafka的acl时使用
	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// 资源名称，和resourceType相关如当resourceType为TOPIC时，则该字段表示topic名称，当resourceType为GROUP时，该字段表示group名称
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 偏移位置
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 个数限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 关键字匹配
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *DescribeACLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeACLRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeACLResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的ACL结果集对象
		Result *AclResponse `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeACLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeACLResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAppInfoRequest struct {
	*tchttp.BaseRequest

	// 偏移位置
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 本次查询用户数目最大数量限制，最大值为50，默认50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAppInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAppInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAppInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的符合要求的App Id列表
		Result *AppIdResponse `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAppInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAppInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConsumerGroupRequest struct {
	*tchttp.BaseRequest

	// ckafka实例id。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 可选，用户需要查询的group名称。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 可选，用户需要查询的group中的对应的topic名称，如果指定了该参数，而group又未指定则忽略该参数。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 本次返回个数限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移位置
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeConsumerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConsumerGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConsumerGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的消费分组信息
		Result *ConsumerGroupResponse `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConsumerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConsumerGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroup struct {

	// groupId
	Group *string `json:"Group,omitempty" name:"Group"`

	// 该 group 使用的协议。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type DescribeGroupInfoRequest struct {
	*tchttp.BaseRequest

	// （过滤条件）按照实例 ID 过滤。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Kafka 消费分组，Consumer-group，这里是数组形式，格式：group.0=xxx&group.1=yyy。
	GroupList []*string `json:"GroupList,omitempty" name:"GroupList" list`
}

func (r *DescribeGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result []*GroupInfoResponse `json:"Result,omitempty" name:"Result" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupOffsetsRequest struct {
	*tchttp.BaseRequest

	// （过滤条件）按照实例 ID 过滤
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Kafka 消费分组
	Group *string `json:"Group,omitempty" name:"Group"`

	// group 订阅的主题名称数组，如果没有该数组，则表示指定的 group 下所有 topic 信息
	Topics []*string `json:"Topics,omitempty" name:"Topics" list`

	// 模糊匹配 topicName
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 本次查询的偏移位置，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 本次返回结果的最大个数，默认为50，最大值为50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeGroupOffsetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupOffsetsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupOffsetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的结果对象
		Result *GroupOffsetResponse `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupOffsetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupOffsetsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 搜索关键字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大返回数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果集列表
		Result []*DescribeGroup `json:"Result,omitempty" name:"Result" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceAttributesRequest struct {
	*tchttp.BaseRequest

	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceAttributesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例属性返回结果对象
		Result *InstanceAttributesResponse `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceAttributesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesDetailRequest struct {
	*tchttp.BaseRequest

	// （过滤条件）按照实例ID过滤
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// （过滤条件）按照实例名称过滤，支持模糊查询
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// （过滤条件）实例的状态。0：创建中，1：运行中，2：删除中，不填默认返回全部
	Status []*int64 `json:"Status,omitempty" name:"Status" list`

	// 偏移量，不填默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，不填则默认10，最大值20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 匹配标签key值。
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 过滤器
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeInstancesDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的实例详情结果对象
		Result *InstanceDetailResponse `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// （过滤条件）按照实例ID过滤
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// （过滤条件）按照实例名称过滤，支持模糊查询
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// （过滤条件）实例的状态。0：创建中，1：运行中，2：删除中，不填默认返回全部
	Status []*int64 `json:"Status,omitempty" name:"Status" list`

	// 偏移量，不填默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，不填则默认10，最大值20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 匹配标签key值。
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的结果
		Result *InstanceResponse `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicAttributesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

func (r *DescribeTopicAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTopicAttributesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的结果对象
		Result *TopicAttributesResponse `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopicAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTopicAttributesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicDetailRequest struct {
	*tchttp.BaseRequest

	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// （过滤条件）按照topicName过滤，支持模糊查询
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 偏移量，不填默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，不填则默认 10，最大值20，取值要大于0
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTopicDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTopicDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的主题详情实体
		Result *TopicDetailResponse `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopicDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTopicDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicRequest struct {
	*tchttp.BaseRequest

	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 过滤条件，按照 topicName 过滤，支持模糊查询
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 偏移量，不填默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTopicRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *TopicResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTopicResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUserRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 按照名称过滤
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 本次返回个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUserRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果列表
		Result *UserResponse `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUserResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 需要过滤的字段。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type Group struct {

	// 组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type GroupInfoMember struct {

	// coordinator 为消费分组中的消费者生成的唯一 ID
	MemberId *string `json:"MemberId,omitempty" name:"MemberId"`

	// 客户消费者 SDK 自己设置的 client.id 信息
	ClientId *string `json:"ClientId,omitempty" name:"ClientId"`

	// 一般存储客户的 IP 地址
	ClientHost *string `json:"ClientHost,omitempty" name:"ClientHost"`

	// 存储着分配给该消费者的 partition 信息
	Assignment *Assignment `json:"Assignment,omitempty" name:"Assignment"`
}

type GroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 错误码，正常为0
		ErrorCode *string `json:"ErrorCode,omitempty" name:"ErrorCode"`

		// group 状态描述（常见的为 Empty、Stable、Dead 三种状态）：
	// Dead：消费分组不存在
	// Empty：消费分组，当前没有任何消费者订阅
	// PreparingRebalance：消费分组处于 rebalance 状态
	// CompletingRebalance：消费分组处于 rebalance 状态
	// Stable：消费分组中各个消费者已经加入，处于稳定状态
		State *string `json:"State,omitempty" name:"State"`

		// 消费分组选择的协议类型正常的消费者一般为 consumer 但有些系统采用了自己的协议如 kafka-connect 用的就是 connect。只有标准的 consumer 协议，本接口才知道具体的分配方式的格式，才能解析到具体的 partition 的分配情况
		ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`

		// 消费者 partition 分配算法常见的有如下几种(Kafka 消费者 SDK 默认的选择项为 range)：range、 roundrobin、 sticky
		Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

		// 仅当 state 为 Stable 且 protocol_type 为 consumer 时， 该数组才包含信息
		Members []*GroupInfoMember `json:"Members,omitempty" name:"Members" list`

		// Kafka 消费分组
		Group *string `json:"Group,omitempty" name:"Group"`
	} `json:"Response"`
}

func (r *GroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GroupInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GroupInfoTopics struct {

	// 分配的 topic 名称
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// 分配的 partition 信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partitions []*int64 `json:"Partitions,omitempty" name:"Partitions" list`
}

type GroupOffsetPartition struct {

	// topic 的 partitionId
	Partition *int64 `json:"Partition,omitempty" name:"Partition"`

	// consumer 提交的 offset 位置
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 支持消费者提交消息时，传入 metadata 作为它用，当前一般为空字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metadata *string `json:"Metadata,omitempty" name:"Metadata"`

	// 错误码
	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// 当前 partition 最新的 offset
	LogEndOffset *int64 `json:"LogEndOffset,omitempty" name:"LogEndOffset"`

	// 未消费的消息个数
	Lag *int64 `json:"Lag,omitempty" name:"Lag"`
}

type GroupOffsetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合调节的总结果数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 该主题分区数组，其中每个元素为一个 json object
	// 注意：此字段可能返回 null，表示取不到有效值。
		TopicList []*GroupOffsetTopic `json:"TopicList,omitempty" name:"TopicList" list`
	} `json:"Response"`
}

func (r *GroupOffsetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GroupOffsetResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GroupOffsetTopic struct {

	// 主题名称
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// 该主题分区数组，其中每个元素为一个 json object
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partitions []*GroupOffsetPartition `json:"Partitions,omitempty" name:"Partitions" list`
}

type Instance struct {

	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例的状态。0：创建中，1：运行中，2：删除中 ， 5 隔离中，-1 创建失败
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 是否开源实例。开源：true，不开源：false
	// 注意：此字段可能返回 null，表示取不到有效值。
	IfCommunity *bool `json:"IfCommunity,omitempty" name:"IfCommunity"`
}

type InstanceAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例ID
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 实例名称
		InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

		// 接入点 VIP 列表信息
		VipList []*VipEntity `json:"VipList,omitempty" name:"VipList" list`

		// 虚拟IP
		Vip *string `json:"Vip,omitempty" name:"Vip"`

		// 虚拟端口
		Vport *string `json:"Vport,omitempty" name:"Vport"`

		// 实例的状态。0：创建中，1：运行中，2：删除中
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 实例带宽，单位：Mbps
		Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

		// 实例的存储大小，单位：GB
		DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

		// 可用区
		ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

		// VPC 的 ID，为空表示是基础网络
		VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

		// 子网 ID， 为空表示基础网络
		SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

		// 实例健康状态， 1：健康，2：告警，3：异常
		Healthy *int64 `json:"Healthy,omitempty" name:"Healthy"`

		// 实例健康信息，当前会展示磁盘利用率，最大长度为256
		HealthyMessage *string `json:"HealthyMessage,omitempty" name:"HealthyMessage"`

		// 创建时间
		CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

		// 消息保存时间,单位为分钟
		MsgRetentionTime *int64 `json:"MsgRetentionTime,omitempty" name:"MsgRetentionTime"`

		// 自动创建 Topic 配置， 若该字段为空，则表示未开启自动创建
		Config *InstanceConfigDO `json:"Config,omitempty" name:"Config"`

		// 剩余创建分区数
		RemainderPartitions *int64 `json:"RemainderPartitions,omitempty" name:"RemainderPartitions"`

		// 剩余创建主题数
		RemainderTopics *int64 `json:"RemainderTopics,omitempty" name:"RemainderTopics"`

		// 当前创建分区数
		CreatedPartitions *int64 `json:"CreatedPartitions,omitempty" name:"CreatedPartitions"`

		// 当前创建主题数
		CreatedTopics *int64 `json:"CreatedTopics,omitempty" name:"CreatedTopics"`

		// 标签数组
	// 注意：此字段可能返回 null，表示取不到有效值。
		Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`

		// 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
		ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
	} `json:"Response"`
}

func (r *InstanceAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InstanceAttributesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InstanceConfigDO struct {

	// 是否自动创建主题
	AutoCreateTopicsEnable *bool `json:"AutoCreateTopicsEnable,omitempty" name:"AutoCreateTopicsEnable"`

	// 分区数
	DefaultNumPartitions *int64 `json:"DefaultNumPartitions,omitempty" name:"DefaultNumPartitions"`

	// 默认的复制Factor
	DefaultReplicationFactor *int64 `json:"DefaultReplicationFactor,omitempty" name:"DefaultReplicationFactor"`
}

type InstanceDetail struct {

	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 访问实例的vip 信息
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 访问实例的端口信息
	Vport *string `json:"Vport,omitempty" name:"Vport"`

	// 虚拟IP列表
	VipList []*VipEntity `json:"VipList,omitempty" name:"VipList" list`

	// 实例的状态。0：创建中，1：运行中，2：删除中：5隔离中， -1 创建失败
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 实例带宽，单位Mbps
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 实例的存储大小，单位GB
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 可用区域ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// vpcId，如果为空，说明是基础网络
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网id
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 实例是否续费，int  枚举值：1表示自动续费，2表示明确不自动续费
	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 实例状态 int：0表示健康，1表示告警，2 表示实例状态异常
	Healthy *string `json:"Healthy,omitempty" name:"Healthy"`

	// 实例状态信息
	HealthyMessage *string `json:"HealthyMessage,omitempty" name:"HealthyMessage"`

	// 实例创建时间时间
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 实例过期时间
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 是否为内部客户。值为1 表示内部客户
	IsInternal *int64 `json:"IsInternal,omitempty" name:"IsInternal"`

	// Topic个数
	TopicNum *int64 `json:"TopicNum,omitempty" name:"TopicNum"`

	// 标识tag
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`
}

type InstanceDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的实例总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 符合条件的实例详情列表
		InstanceList []*InstanceDetail `json:"InstanceList,omitempty" name:"InstanceList" list`
	} `json:"Response"`
}

func (r *InstanceDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InstanceDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		InstanceList []*Instance `json:"InstanceList,omitempty" name:"InstanceList" list`

		// 符合条件的结果总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	} `json:"Response"`
}

func (r *InstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type JgwOperateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的code，0为正常，非0为错误
		ReturnCode *string `json:"ReturnCode,omitempty" name:"ReturnCode"`

		// 成功消息
		ReturnMessage *string `json:"ReturnMessage,omitempty" name:"ReturnMessage"`
	} `json:"Response"`
}

func (r *JgwOperateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *JgwOperateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyGroupOffsetsRequest struct {
	*tchttp.BaseRequest

	// kafka实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// kafka 消费分组
	Group *string `json:"Group,omitempty" name:"Group"`

	// 重置offset的策略，入参含义 0. 对齐shift-by参数，代表把offset向前或向后移动shift条 1. 对齐参考(by-duration,to-datetime,to-earliest,to-latest),代表把offset移动到指定timestamp的位置 2. 对齐参考(to-offset)，代表把offset移动到指定的offset位置
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`

	// 表示需要重置的topics， 不填表示全部
	Topics []*string `json:"Topics,omitempty" name:"Topics" list`

	// 当strategy为0时，必须包含该字段，可以大于零代表会把offset向后移动shift条，小于零则将offset向前回溯shift条数。正确重置后新的offset应该是(old_offset + shift)，需要注意的是如果新的offset小于partition的earliest则会设置为earliest，如果大于partition 的latest则会设置为latest
	Shift *int64 `json:"Shift,omitempty" name:"Shift"`

	// 单位ms。当strategy为1时，必须包含该字段，其中-2表示重置offset到最开始的位置，-1表示重置到最新的位置(相当于清空)，其它值则代表指定的时间，会获取topic中指定时间的offset然后进行重置，需要注意的时，如果指定的时间不存在消息，则获取最末尾的offset。
	ShiftTimestamp *int64 `json:"ShiftTimestamp,omitempty" name:"ShiftTimestamp"`

	// 需要重新设置的offset位置。当strategy为2，必须包含该字段。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *ModifyGroupOffsetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyGroupOffsetsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyGroupOffsetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果
		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyGroupOffsetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyGroupOffsetsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceAttributesConfig struct {

	// 自动创建 true 表示开启，false 表示不开启
	AutoCreateTopicEnable *bool `json:"AutoCreateTopicEnable,omitempty" name:"AutoCreateTopicEnable"`

	// 可选，如果auto.create.topic.enable设置为true没有设置该值时，默认设置为3
	DefaultNumPartitions *int64 `json:"DefaultNumPartitions,omitempty" name:"DefaultNumPartitions"`

	// 如歌auto.create.topic.enable设置为true没有指定该值时默认设置为2
	DefaultReplicationFactor *int64 `json:"DefaultReplicationFactor,omitempty" name:"DefaultReplicationFactor"`
}

type ModifyInstanceAttributesRequest struct {
	*tchttp.BaseRequest

	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例日志的最长保留时间，单位分钟，最大30天，0代表不开启日志保留时间回收策略
	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitempty" name:"MsgRetentionTime"`

	// 实例名称，是一个不超过 64 个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例配置
	Config *ModifyInstanceAttributesConfig `json:"Config,omitempty" name:"Config"`
}

func (r *ModifyInstanceAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceAttributesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果
		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceAttributesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPasswordRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用户名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用户当前密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 用户新密码
	PasswordNew *string `json:"PasswordNew,omitempty" name:"PasswordNew"`
}

func (r *ModifyPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPasswordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果
		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPasswordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTopicAttributesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主题名称。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 主题备注，是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线-。
	Note *string `json:"Note,omitempty" name:"Note"`

	// IP 白名单开关，1：打开；0：关闭。
	EnableWhiteList *int64 `json:"EnableWhiteList,omitempty" name:"EnableWhiteList"`

	// 默认为1。
	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitempty" name:"MinInsyncReplicas"`

	// 默认为 0，0：false；1：true。
	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitempty" name:"UncleanLeaderElectionEnable"`

	// 消息保留时间，单位：ms，当前最小值为60000ms。
	RetentionMs *int64 `json:"RetentionMs,omitempty" name:"RetentionMs"`

	// Segment 分片滚动的时长，单位：ms，当前最小为86400000ms。
	SegmentMs *int64 `json:"SegmentMs,omitempty" name:"SegmentMs"`

	// 主题消息最大值，单位为 Byte，最大值为8388608Byte（即8MB）。
	MaxMessageBytes *int64 `json:"MaxMessageBytes,omitempty" name:"MaxMessageBytes"`

	// 消息删除策略，可以选择delete 或者compact
	CleanUpPolicy *string `json:"CleanUpPolicy,omitempty" name:"CleanUpPolicy"`
}

func (r *ModifyTopicAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTopicAttributesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTopicAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果集
		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTopicAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTopicAttributesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Partition struct {

	// 分区ID
	PartitionId *int64 `json:"PartitionId,omitempty" name:"PartitionId"`
}

type SubscribedInfo struct {

	// 订阅的主题名
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 订阅的分区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partition []*int64 `json:"Partition,omitempty" name:"Partition" list`
}

type Tag struct {

	// 标识的key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标识的值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type Topic struct {

	// 主题的ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 主题的名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Note *string `json:"Note,omitempty" name:"Note"`
}

type TopicAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 主题 ID
		TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

		// 创建时间
		CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

		// 主题备注
	// 注意：此字段可能返回 null，表示取不到有效值。
		Note *string `json:"Note,omitempty" name:"Note"`

		// 分区个数
		PartitionNum *int64 `json:"PartitionNum,omitempty" name:"PartitionNum"`

		// IP 白名单开关，1：打开； 0：关闭
		EnableWhiteList *int64 `json:"EnableWhiteList,omitempty" name:"EnableWhiteList"`

		// IP 白名单列表
		IpWhiteList []*string `json:"IpWhiteList,omitempty" name:"IpWhiteList" list`

		// topic 配置数组
		Config *Config `json:"Config,omitempty" name:"Config"`

		// 分区详情
		Partitions []*TopicPartitionDO `json:"Partitions,omitempty" name:"Partitions" list`
	} `json:"Response"`
}

func (r *TopicAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TopicAttributesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TopicDetail struct {

	// 主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 分区数
	PartitionNum *int64 `json:"PartitionNum,omitempty" name:"PartitionNum"`

	// 副本数
	ReplicaNum *int64 `json:"ReplicaNum,omitempty" name:"ReplicaNum"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Note *string `json:"Note,omitempty" name:"Note"`

	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 是否开启ip鉴权白名单，true表示开启，false表示不开启
	EnableWhiteList *bool `json:"EnableWhiteList,omitempty" name:"EnableWhiteList"`

	// ip白名单中ip个数
	IpWhiteListCount *int64 `json:"IpWhiteListCount,omitempty" name:"IpWhiteListCount"`

	// 数据备份cos bucket: 转存到cos 的bucket地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForwardCosBucket *string `json:"ForwardCosBucket,omitempty" name:"ForwardCosBucket"`

	// 数据备份cos 状态： 1 不开启数据备份，0 开启数据备份
	ForwardStatus *int64 `json:"ForwardStatus,omitempty" name:"ForwardStatus"`

	// 数据备份到cos的周期频率
	ForwardInterval *int64 `json:"ForwardInterval,omitempty" name:"ForwardInterval"`

	// 高级配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Config *Config `json:"Config,omitempty" name:"Config"`
}

type TopicDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的主题详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		TopicList []*TopicDetail `json:"TopicList,omitempty" name:"TopicList" list`

		// 符合条件的所有主题详情数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	} `json:"Response"`
}

func (r *TopicDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TopicDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TopicPartitionDO struct {

	// Partition ID
	Partition *int64 `json:"Partition,omitempty" name:"Partition"`

	// Leader 运行状态
	LeaderStatus *int64 `json:"LeaderStatus,omitempty" name:"LeaderStatus"`

	// ISR 个数
	IsrNum *int64 `json:"IsrNum,omitempty" name:"IsrNum"`

	// 副本个数
	ReplicaNum *int64 `json:"ReplicaNum,omitempty" name:"ReplicaNum"`
}

type TopicResult struct {

	// 返回的主题信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicList []*Topic `json:"TopicList,omitempty" name:"TopicList" list`

	// 符合条件的 topic 数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type User struct {

	// 用户id
	UserId *int64 `json:"UserId,omitempty" name:"UserId"`

	// 用户名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 最后更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type UserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的用户列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Users []*User `json:"Users,omitempty" name:"Users" list`

		// 符合条件的总用户数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	} `json:"Response"`
}

func (r *UserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UserResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VipEntity struct {

	// 虚拟IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 虚拟端口
	Vport *string `json:"Vport,omitempty" name:"Vport"`
}
