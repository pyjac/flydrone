import * as types from './action_types'
import { moveDrone, markDroneInactive } from './actions'
import mqtt from 'mqtt'


const droneTimers = []

const droneMove = (dispatch, payload) => {
  return new Promise((resolve, reject) => {
    dispatch(moveDrone(payload))
    resolve()
  })
}

const markDroneInactiveTimer = (dispatch, droneId) => {
  return setTimeout(() => dispatch(markDroneInactive(droneId)), 10000);
}
const setupDroneClient = (dispatch) => {
  const client = mqtt.connect(process.env.REACT_APP_SOCKET_URL) // you add a ws:// url here
  client.subscribe(types.DRONE_MOVE)
  client.on("message", function (topic, payload) {
    switch (topic) {
      case types.DRONE_MOVE:
        const dronePayload = JSON.parse(new TextDecoder("utf-8").decode(payload));
        droneMove(dispatch, dronePayload).then(() => {
          const droneId = dronePayload.id
          let droneTimer = droneTimers[droneId]
          if (droneTimer) {
            clearTimeout(droneTimer)
          }
          droneTimers[droneId] = markDroneInactiveTimer(dispatch, droneId)
        });
        break
      default:
        break
    }
  })

  return client
}

export default setupDroneClient