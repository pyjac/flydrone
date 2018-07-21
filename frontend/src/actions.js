import * as types from './action_types'

export const moveDrone = drone => ({
  type: types.DRONE_MOVE,
  drone
})

export const markDroneInactive = droneId => ({
  type: types.MARK_DRONE_INACTIVE,
  droneId
})