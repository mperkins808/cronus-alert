import axios, { AxiosResponse } from "axios";

export type Alert = {
  summary: String;
  description: String;
  status?: Status;
};

export enum Status {
  Firing = "firing",
  Resolved = "resolved",
}

export class CronusAlert {
  private authToken: string;

  constructor(authToken: string) {
    this.authToken = authToken;
  }

  private buildAlert(alert: Alert) {
    var innerAlert = {
      status: "Generic",
      annotations: {
        description: alert.description,
        summary: alert.summary,
      },
    };
    if (alert.status) {
      innerAlert.status = alert.status;
    }
    var outerAlert = {
      alerts: [innerAlert],
    };
    return outerAlert;
  }

  async FireAlert(
    summary: string,
    description: string,
    status?: Status
  ): Promise<AxiosResponse<any>> {
    const apiUrl = "https://cronusmonitoring.com/api/alert";

    const alert: Alert = {
      summary: summary,
      description: description,
      status: status,
    };

    const jsonAlert = this.buildAlert(alert);

    try {
      const response = await axios.post(apiUrl, jsonAlert, {
        headers: {
          Authorization: `Bearer ${this.authToken}`,
          "Content-Type": "application/json",
        },
      });

      return response;
    } catch (error: any) {
      if (error.response) {
        throw error.response;
      }

      throw error;
    }
  }
}
