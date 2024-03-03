# CapitolReef - Excel -> ICS

The Vacation Planner Calendar Importer is a Go-based tool designed to streamline the process of creating detailed vacation itineraries in excel and importing them into Google Calendar. By parsing event details from a spreadsheet, formatting them for Google Calendar compatibility, and generating an iCalendar (.ics) file, this tool simplifies the task of planning and organizing vacation activities.

## Features

- **Spreadsheet Parsing**: Converts event details from a spreadsheet into a structured format.
- **ChatGPT Integration**: Utilizes OpenAI's ChatGPT to format event details suitable for calendar entries. Allows for generic spreadsheets to work.
- **ICS Generation**: Creates a universally compatible `.ics` file that can be imported into Google Calendar and other calendar applications.
- **Timezone Support**: Ensures that event times are correctly aligned with their respective timezones. Done automatically based on location (guessed by activites)

## Prerequisites

Before you start using the Vacation Planner Calendar Importer, ensure you have the following installed:

- Go (version 1.15 or newer)
- Access to OpenAI's API (ChatGPT)

## Installation

1. **Clone the Repository**:


2. **Install Dependencies**:
Navigate to the project directory and run: 
 ```go mod tidy ```

## Usage

1. **Prepare Your Spreadsheet**:
- Format your vacation event details in a spreadsheet with columns for date, start time, end time, summary, description, and location.
- Save the spreadsheet in a location accessible by the tool.

2. **Set Environment Variables**:
- Ensure your OpenAI API key is set as an environment variable:
  ```
  export OPENAI_API_KEY='your_openai_api_key_here'
  ```

3. **Run the Tool**:
- Execute the main program, specifying the path to your spreadsheet:
  ```
  go run cmd/main.go --file path/to/your/spreadsheet.xlsx
  ```

4. **Generate ICS File**:
- The tool parses the spreadsheet, formats the event details, and generates an `.ics` file in the project directory.

5. **Import into Google Calendar**:
- Open Google Calendar and navigate to `Settings > Import & Export`.
- Click on `Select file from your computer` and choose the generated `.ics` file.
- Select the calendar where you want to import the events and click `Import`.

## Customization

- You can modify the `parser` and `ics` packages to adjust the event detail extraction logic and the `.ics` file formatting to suit your specific needs.

## Troubleshooting

- **API Key Issues**: Ensure the OPENAI_API_KEY environment variable is correctly set.
- **Spreadsheet Format**: Verify that the spreadsheet follows the expected format.
- **ICS Import Errors**: Check the `.ics` file for any formatting issues if Google Calendar rejects the import.

## Contributing

Contributions to the Vacation Planner Calendar Importer are welcome. Please feel free to fork the repository, make your changes, and submit a pull request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
