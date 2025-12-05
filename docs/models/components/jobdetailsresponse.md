# JobDetailsResponse


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `JobID`                                                      | *string*                                                     | :heavy_check_mark:                                           | The unique identifier of the job.                            |
| `JobCode`                                                    | *string*                                                     | :heavy_check_mark:                                           | The code representing the type of job executed.              |
| `Status`                                                     | [components.JobStatus](../../models/components/jobstatus.md) | :heavy_check_mark:                                           | The status of the job executed.                              |
| `Output`                                                     | **string*                                                    | :heavy_minus_sign:                                           | The status of the job executed.                              |
| `SourceID`                                                   | *string*                                                     | :heavy_check_mark:                                           | The identifier of the source against which the job executed. |
| `SourceName`                                                 | *string*                                                     | :heavy_check_mark:                                           | The name of the source against which the job executed.       |
| `CreatedAt`                                                  | *int64*                                                      | :heavy_check_mark:                                           | The time the job execution was created.                      |
| `UpdatedAt`                                                  | *int64*                                                      | :heavy_check_mark:                                           | The time the job execution was updated.                      |