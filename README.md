# Dynafetch

Dynafetch is a CLI tool for testing FetchXML queries on a Microsoft Dynamics server.  
It allows you to experiment with queries by automatically fetching results upon changes.

## Installation

Install Dynafetch using:

```sh
go install <path-to-dynafetch>
```

## Usage

Run Dynafetch with the path to your FetchXML file:

```sh
dynafetch path/to/fetch.xml
```

The tool waits for credentials, which are received from a Chrome extension.  
Once authenticated, it sends a request to the Dynamics server and prints the retrieved records to the terminal.  
Dynafetch also watches for file changes, re-executing the query and printing the updated results.

### Example usage:

```sh
dynafetch queries/accounts.xml
```

Output:

```
scheduledstart 2025-05-13T09:00:00Z
Meeting with colleagues #96
isalldayevent false
activityid 20f2ce7a-b1e3-ef11-9341-7c1e5251d46e
activityadditionalparams {"scheduledstartformatted":"5/13/2025 11:00 AM"}


activityid 22f2ce7a-b1e3-ef11-9341-7c1e5251d46e
activityadditionalparams {"scheduledstartformatted":"5/16/2025 11:00 AM"}
scheduledstart 2025-05-16T09:00:00Z
Personal meeting #97
isalldayevent false
```

After modifying `queries/accounts.xml`:

```
scheduledstart 2025-05-13T09:00:00Z
subject Meeting with colleagues #96
isalldayevent false
activityid 20f2ce7a-b1e3-ef11-9341-7c1e5251d46e
activityadditionalparams {"scheduledstartformatted":"5/13/2025 11:00 AM"}


activityid 22f2ce7a-b1e3-ef11-9341-7c1e5251d46e
activityadditionalparams {"scheduledstartformatted":"5/16/2025 11:00 AM"}
scheduledstart 2025-05-16T09:00:00Z
subject Personal meeting #97
isalldayevent false


scheduledstart 2025-05-19T09:00:00Z
subject Videocall #98
isalldayevent false
activityid 24f2ce7a-b1e3-ef11-9341-7c1e5251d46e
activityadditionalparams {"scheduledstartformatted":"5/19/2025 11:00 AM"}
```

## License

This project is licensed under the MIT License.

## Contributing

Pull requests are welcome! Feel free to open an issue or submit improvements.
