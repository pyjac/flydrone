import * as types from './action_types'

export const moveDrone = drone => ({
  type: types.DRONE_MOVE,
  drone
})