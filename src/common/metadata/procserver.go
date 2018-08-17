/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */
package metadata

import (
	"time"
)

type ProcModuleConfig struct {
	ApplicationID int64  `json:"bk_biz_id"`
	ModuleName    string `json:"bk_module_name"`
	ProcessID     int64  `json:"bk_process_id"`
}

type ProcInstanceModel struct {
	ApplicationID  int64  `json:"bk_biz_id" bson:"bk_biz_id"`
	SetID          int64  `json:"bk_set_id" bson:"bk_set_id,omitempty"`
	ModuleID       int64  `json:"bk_module_id" bson:"bk_module_id,omitempty"`
	ProcID         int64  `json:"bk_process_id" bson:"bk_process_id"`
	FuncID         int64  `json:"bk_func_id" bson:"bk_func_id"`
	ProcInstanceID uint64 `json:"proc_instance_id" bson:"proc_instance_id"`
	HostID         int64  `json:"bk_host_id" bson:"bk_host_id"`
	HostInstanID   uint64 `json:"bk_host_instance_id" bson:"bk_host_instance_id"`
	HostProcID     uint64 `json:"host_proc_id" bson:"host_proc_id"`
	OwnerID        string `json:"bk_supplier_account" bson:"bk_supplier_account"`
}

type ProcessOperate struct {
	ApplicationID  int64  `json:"bk_biz_id"`
	SetName        string `json:"bk_set_name"`
	ModuleName     string `json:"bk_module_name"`
	FuncID         string `json:"bk_func_id"`
	HostInstanceID string `json:"bk_host_instance_id"`
	OpType         int    `json:"bk_proc_optype"`
}

type ProcModuleResult struct {
	BaseResp `json:",inline"`
	Data     []ProcModuleConfig `json:"data"`
}

type ProcInstModelResult struct {
	BaseResp `json:",inline"`
	Data     struct {
		Count int                 `json:"count"`
		Info  []ProcInstanceModel `json:"info"`
	} `json:"data"`
}

type GseHost struct {
	HostID       int64
	Ip           string `json:"ip,omitempty"`
	BkCloudId    int64  `json:"bk_cloud_id,omitempty"`
	BkSupplierId int64  `json:"bk_supplier_ed,omitempty"`
}

type GseProcMeta struct {
	Namespace string            `json:"namespace,omitempty"`
	Name      string            `json:"name,omitempty"`
	Labels    map[string]string `json:"labels,omitempty"`
}

type GseProcRequest struct {
	AppID    int64       `json:"bk_biz_id"  bson:"bk_biz_id"`
	ModuleID int64       `json:"bk_module_id" bson:"bk_module_id"`
	ProcID   int64       `json:"bk_process_id" bson:"bk_process_id"`
	Meta     GseProcMeta `json:"meta,omitempty" bson:"meta"`
	Hosts    []GseHost   `json:"hosts,omitempty" bson:"hosts"`
	OpType   int         `json:"op_type,omitempty" bson:"-"`
	Spec     GseProcSpec `json:"spec,omitempty" bson:"spec"`
}

type ProcInstanceDetail struct {
	GseProcRequest `json:",inline"`
	OwnerID        string                   `json:"bk_supplier_account" bson:"bk_supplier_account"`
	HostID         int64                    `json:"bk_hsot_id", bson:"bk_host_id"`
	Status         ProcInstanceDetailStatus `json:"status" bson:"status"` //1 register gse sucess, 2 register error need retry 3 unregister error need retry
}

type ProcInstanceDetailResult struct {
	BaseResp `json:",inline"`

	Data struct {
		Count int                  `json:"count"`
		Info  []ProcInstanceDetail `json:"info"`
	} `json:"data"`
}

type ProcInstanceDetailStatus int64

var (
	ProcInstanceDetailStatusRegisterSucc     ProcInstanceDetailStatus = 1
	ProcInstanceDetailStatusRegisterFailed   ProcInstanceDetailStatus = 2
	ProcInstanceDetailStatusUnRegisterFailed ProcInstanceDetailStatus = 10
)

type ModifyProcInstanceDetail struct {
	Conds map[string]interface{} `json:"condition"`
	Data  map[string]interface{} `json:"data"`
}

type GseProcSpec struct {
	Identity         GseProcIdentity         `json:"identity,omitempty"`
	Control          GseProcControl          `json:"control,omitempty"`
	Resource         GseProcResource         `json:"resource,omitempty"`
	MonitorPolicy    GseProcMonitorPlolicy   `json:"monitor_policy,omitempty"`
	WarnReportPolicy GseProcWarnReportPolicy `json:"warn_report_policy,omitempty"`
	Configmap        GseProcConfigmap        `json:"configmap,omitempty"`
}

type GseProcIdentity struct {
	IndexKey   string `json:"index_key,omitempty"`
	ProcName   string `json:"proc_name,omitempty"`
	SetupPath  string `json:"setup_path,omitempty"`
	PidPath    string `json:"pid_path,omitempty"`
	ConfigPath string `json:"config_path,omitempty"`
	LogPath    string `json:"log_path,omitempty"`
}

type GseProcControl struct {
	StartCmd   string `json:"start_cmd,omitempty"`
	StopCmd    string `json:"stop_cmd,omitempty"`
	RestartCmd string `json:"restart_cmd,omitempty"`
	ReloadCmd  string `json:"reload_cmd,omitempty"`
	KillCmd    string `json:"kill_cmd,omitempty"`
	VersionCmd string `json:"version_cmd,omitempty"`
	HealthCmd  string `json:"health_cmd,omitempty"`
}

type GseProcResource struct {
	Cpu  float64 `json:"cpu,omitempty"`
	Mem  float64 `json:"mem,omitempty"`
	Fd   int     `json:"fd,omitempty"`
	Disk int     `json:"disk,omitempty"`
	Net  int     `json:"net,omitempty"`
}

type GseProcMonitorPlolicy struct {
	AutoType       int    `json:"auto_type,omitempty"`
	StartCheckSecs int    `json:"start_check_secs,omitempty"`
	StopCheckSecs  int    `json:"stop_check_secs,omitempty"`
	StartRetries   int    `json:"start_retries,omitempty"`
	StartInterval  int    `json:"start_interval,omitempty"`
	CrontabRule    string `json:"crontab_rule,omitempty"`
}

type GseProcWarnReportPolicy struct {
	ReportId int `json:"report_id,omitempty"`
}

type GseProcConfigmap struct {
	Path string `json:"path,omitempty"`
	Md5  string `json:"md5,omitempty"`
}

type FilePriviewMap struct {
	Content string `json:"content"`
	Inst    string `json:"inst"`
}

// InlineProcInfo process info convert gse proc info
type InlineProcInfo struct {
	//Meta    GseProcMeta
	//Spec    GseProcSpec
	ProcInfo map[string]interface{}
	ProcNum  int64
	AppID    int64 // use gse proc namespace
	FunID    int64
}

type ProcessOperateTask struct {
	OperateInfo *ProcessOperate                     `json:"operate_info" bson:"operate_info"`
	TaskID      string                              `json:"task_id" bson:"task_id"`
	GseTaskID   string                              `json:"gse_task_id" bson:"gse_task_id"`
	Namespace   string                              `json:"namespace" bson:"namespace"`
	Status      ProcessOperateTaskStatus            `json:"status" bson:"status"`
	CreateTime  time.Time                           `json:"create_time" bson:"create_time"`
	OwnerID     string                              `json:"bk_supplier_account" bson:"bk_supplier_account"`
	User        string                              `json:"user,omitempty" bson:"user,omitempty"`
	Detail      map[string]ProcessOperateTaskDetail `json:"detail" bson:"detail"`
}

type ProcessOperateTaskStatus int64

var (
	ProcessOperateTaskStatusWaitOP ProcessOperateTaskStatus = 1
	ProcessOperateTaskStatusRuning ProcessOperateTaskStatus = 115
	ProcessOperateTaskStatusErr    ProcessOperateTaskStatus = 2
	ProcessOperateTaskStatusSucc   ProcessOperateTaskStatus = 3
)

type ProcessOperateTaskResult struct {
	BaseResp `json:",inline"`

	Data struct {
		Count int                  `json:"count"`
		Info  []ProcessOperateTask `json:"info"`
	} `json:"data"`
}

type ProcessOperateTaskDetail struct {
	Errcode int    `json:"errcode" bson:"error_code"`
	ErrMsg  string `json:"errmsg" bson:"error_msg"`
}

type GseProcResponeBase struct {
	Code   int    `json:"bk_error_code"`
	ErrMsg string `json:"bk_error_msg"`
}
type GseProcessOperateTaskResult struct {
	Result             map[string]ProcessOperateTaskDetail `json:"result"`
	GseProcResponeBase `json:",inline"`
}

type GseProcRespone struct {
	GseProcResponeBase `json:",inline"`
	Result             map[string]interface{} `json:"result"`
}
