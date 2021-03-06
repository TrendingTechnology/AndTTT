// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: response.proto

package com.github.dawidd6.andttt.proto;

public interface ResponseOrBuilder extends
    // @@protoc_insertion_point(interface_extends:proto.Response)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>.proto.Error error = 1000;</code>
   */
  int getErrorValue();
  /**
   * <code>.proto.Error error = 1000;</code>
   */
  com.github.dawidd6.andttt.proto.Error getError();

  /**
   * <code>.proto.UnrecognizedResponse unrecognized = 9999;</code>
   */
  com.github.dawidd6.andttt.proto.UnrecognizedResponse getUnrecognized();
  /**
   * <code>.proto.UnrecognizedResponse unrecognized = 9999;</code>
   */
  com.github.dawidd6.andttt.proto.UnrecognizedResponseOrBuilder getUnrecognizedOrBuilder();

  /**
   * <code>.proto.RegisterNameResponse register_name = 1;</code>
   */
  com.github.dawidd6.andttt.proto.RegisterNameResponse getRegisterName();
  /**
   * <code>.proto.RegisterNameResponse register_name = 1;</code>
   */
  com.github.dawidd6.andttt.proto.RegisterNameResponseOrBuilder getRegisterNameOrBuilder();

  /**
   * <code>.proto.CreateRoomResponse create_room = 2;</code>
   */
  com.github.dawidd6.andttt.proto.CreateRoomResponse getCreateRoom();
  /**
   * <code>.proto.CreateRoomResponse create_room = 2;</code>
   */
  com.github.dawidd6.andttt.proto.CreateRoomResponseOrBuilder getCreateRoomOrBuilder();

  /**
   * <code>.proto.JoinRoomResponse join_room = 3;</code>
   */
  com.github.dawidd6.andttt.proto.JoinRoomResponse getJoinRoom();
  /**
   * <code>.proto.JoinRoomResponse join_room = 3;</code>
   */
  com.github.dawidd6.andttt.proto.JoinRoomResponseOrBuilder getJoinRoomOrBuilder();

  /**
   * <code>.proto.LeaveRoomResponse leave_room = 4;</code>
   */
  com.github.dawidd6.andttt.proto.LeaveRoomResponse getLeaveRoom();
  /**
   * <code>.proto.LeaveRoomResponse leave_room = 4;</code>
   */
  com.github.dawidd6.andttt.proto.LeaveRoomResponseOrBuilder getLeaveRoomOrBuilder();

  /**
   * <code>.proto.GetRoomsResponse get_rooms = 5;</code>
   */
  com.github.dawidd6.andttt.proto.GetRoomsResponse getGetRooms();
  /**
   * <code>.proto.GetRoomsResponse get_rooms = 5;</code>
   */
  com.github.dawidd6.andttt.proto.GetRoomsResponseOrBuilder getGetRoomsOrBuilder();

  /**
   * <code>.proto.MoveResponse move = 6;</code>
   */
  com.github.dawidd6.andttt.proto.MoveResponse getMove();
  /**
   * <code>.proto.MoveResponse move = 6;</code>
   */
  com.github.dawidd6.andttt.proto.MoveResponseOrBuilder getMoveOrBuilder();

  /**
   * <code>.proto.RestartResponse restart = 7;</code>
   */
  com.github.dawidd6.andttt.proto.RestartResponse getRestart();
  /**
   * <code>.proto.RestartResponse restart = 7;</code>
   */
  com.github.dawidd6.andttt.proto.RestartResponseOrBuilder getRestartOrBuilder();

  /**
   * <code>.proto.StarterPackResponse starter_pack = 8;</code>
   */
  com.github.dawidd6.andttt.proto.StarterPackResponse getStarterPack();
  /**
   * <code>.proto.StarterPackResponse starter_pack = 8;</code>
   */
  com.github.dawidd6.andttt.proto.StarterPackResponseOrBuilder getStarterPackOrBuilder();

  /**
   * <code>.proto.EnemyDisconnectedResponse enemy_disconnected = 101;</code>
   */
  com.github.dawidd6.andttt.proto.EnemyDisconnectedResponse getEnemyDisconnected();
  /**
   * <code>.proto.EnemyDisconnectedResponse enemy_disconnected = 101;</code>
   */
  com.github.dawidd6.andttt.proto.EnemyDisconnectedResponseOrBuilder getEnemyDisconnectedOrBuilder();

  /**
   * <code>.proto.EnemyLeftResponse enemy_left = 102;</code>
   */
  com.github.dawidd6.andttt.proto.EnemyLeftResponse getEnemyLeft();
  /**
   * <code>.proto.EnemyLeftResponse enemy_left = 102;</code>
   */
  com.github.dawidd6.andttt.proto.EnemyLeftResponseOrBuilder getEnemyLeftOrBuilder();

  /**
   * <code>.proto.EnemyMovedResponse enemy_moved = 103;</code>
   */
  com.github.dawidd6.andttt.proto.EnemyMovedResponse getEnemyMoved();
  /**
   * <code>.proto.EnemyMovedResponse enemy_moved = 103;</code>
   */
  com.github.dawidd6.andttt.proto.EnemyMovedResponseOrBuilder getEnemyMovedOrBuilder();

  public com.github.dawidd6.andttt.proto.Response.TypeCase getTypeCase();
}
