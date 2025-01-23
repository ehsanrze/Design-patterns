import {Notification} from "./Notification";


export class PatientVerifiedNotification extends Notification {
    send(recipient: string, inputs: {
        patientId: number
    }): void {
        this.implementor.sendNotification(recipient, `You patient with ID ${inputs.patientId} has rejected.`)
    }
}