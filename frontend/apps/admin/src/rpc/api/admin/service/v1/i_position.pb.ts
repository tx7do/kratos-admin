// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.0
//   protoc               unknown
// source: admin/service/v1/i_position.proto

/* eslint-disable */
import { type Empty } from "../../../google/protobuf/empty.pb";
import { type PagingRequest } from "../../../pagination/v1/pagination.pb";
import {
  type CreatePositionRequest,
  type DeletePositionRequest,
  type GetPositionRequest,
  type ListPositionResponse,
  type Position,
  type UpdatePositionRequest,
} from "../../../user/service/v1/position.pb";

/** 职位管理服务 */
export interface PositionService {
  /** 查询职位列表 */
  ListPosition(request: PagingRequest): Promise<ListPositionResponse>;
  /** 查询职位详情 */
  GetPosition(request: GetPositionRequest): Promise<Position>;
  /** 创建职位 */
  CreatePosition(request: CreatePositionRequest): Promise<Empty>;
  /** 更新职位 */
  UpdatePosition(request: UpdatePositionRequest): Promise<Empty>;
  /** 删除职位 */
  DeletePosition(request: DeletePositionRequest): Promise<Empty>;
}
