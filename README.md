 
# App Setting Finder (diglett) - **Abandoned**

This application (early development) is intended to inspect all application code in order to map all app_settings.
The intention is to document all app_settings found and describe the purpose of each one of them.

This application is being developed using the GO Programming language (1.18).


## Application folder structure

    .
    ├── internal               # This package holds the private library code used in your application. Should not be shared with other services.
    │   ├── menus              # Folder where all files referring to the application menu will be allocated
    │   ├── messages           # Folder containing the file with all the messages used in the application  
    ├── pkg                    # This folder contains code which is OK for other services to consume.
    │   ├── logger             
    │   └── report
    ├── input                  # This folder holds the files containing the app_settings exported from the database    
    ├── output                 # This folder holds the files containing the execution result
    ├── logs                   # Logs generated during execution
    ├── tests                  # Unit tests
    ├── configs                # This folder holds the config files
    ├── go.mod                 # The go. mod file defines the module's module path, which is also the import path used for the root directory, and its dependency requirements, which are the other modules needed for a successful build.      
    ├── main.go
    └── README.md
## Expected features

| Reference | Description               | Status                                                |
| --- | ----------------- | ---------------------------------------------------------------- |
| REF001 | Create a menu containing the options described in the references *(REF001)*. | In progress |
| REF002 | Search for app_settings considering the different ways (AppSettingManager, AppSetting::model()->get ...). |  In progress |
| REF003 | Generate a file containing the successfully identified app_settings. | Done |
| REF004 | Generate a file containing the app_settings that needs manual check. | Done |
| REF005 | Write log files (an external lib can be used in the future, which would provide better performance).| Done |
| REF006 | Compare app_settings found during analysis versus app_settings found in a UAT database.| Pending |
| REF007 | Automatically categorize an app_setting.| Pending |
| REF008 | Export categorized app_settings to a xlsx file.| Pending |
| REF009 | Create a local database, in order to disregard the app_settings that requires manual verification.| Pending |
| REF010 | Create a web interface that allows the developer to handle app_settings that requires manual verification.| Pending |
| REF011 | Adapt application folder structure allowing building more than one app.| Pending |
| REF012 | Provide executables compatible with Linux, Mac OS and Windows.| Pending |

### References

**REF001** - Create a menu with the following options:

```
1 - Setup
2 - Search for App Settings
3 - Compare App Settings
4 - Start web service
5 - Close application
```

- **Setup** - This option will allow the development to set the path of the App folder (where contains the php files). 
     - The value informed by the developer must be saved in a file placed in the config folder.

**Configuration file example**
```
APP_PATH=/home/user/PHPApp
```
- **Search for App Settings** - This option should trigger the search for app_settings.

- **Compare App Settings** - This option should trigger a comparison for app_settings. For more details read the reference *(REF006)*.

- **Start web service** - This option should launch the web app, allowing the developer to manipulate the app_settings that need manual handling. For more details read the reference *(REF010)*.

- **Close application** - Closes the application.

--- 

**REF002** - Write a routine that automatically searches for app_settings in the source code of the PHP application.\
This routine should consider all possible ways to use an app_setting, *for example*:
```
- AppSettingManager::get('app_setting.name')
- AppSetting::model()->findByAttributes(array('option' => 'app_setting.name'))
```
**Note:** *I believe we can read the file all at once rather than line by line, considering the files aren't that big*.

---

**REF003** - Generate a simple file to insert the app_settings that were successfully found.\
This file must be placed inside the output folder and must always be regenerated every time the application is executed.

---
**REF004** - Generate a simple file to insert the app_settings were found but needs manual checking.\
This file must be placed inside the output folder and must always be regenerated every time the application is executed.

---
**REF005** - Create a feature as simple as possible that allows the developer to add logs.\
This feature must provide at least four levels of logging (**Info**, **Debug**, **Warn** and **Error**).\
The log file must always be regenerated every time the application is executed.

---
**REF006** - This functionality should compare the output.txt file versus the input.csv file (for example).\
The result of this comparison should print which app_setting exists in the database and does not exist in the source code and vice versa.\
At the end of the comparison process, a report should be printed on the screen, informing the number of differences found.\
A file should also be saved inside the output folder, called comparison.txt where all the differences found will be add.

---
**REF007 / REF008** - Create a feature to categorize the app_setting during the analysis process. This feature should take into account the points below:

- Add a slice with keywords
- A new file named *output_categorized.xlsx* must be created inside the output folder. All categorized app_settings should be add in this file.
- The algorithm should check if the app_setting content contains any of the defined keywords, for example:
```
App setting = feauture.google.enable
```
In the case above, the algorithm should add the category = google.

- The second option refers to the file path. Files are often organized within the project folder.

- The third and last check refers to the file name, if the file name has a keyword, then that category can be used.

**Note**: If we found a keyword in the first check, then the other checks can be disregarded.

---

**REF009** - Analyze best database solution. Maybe sqlite?.
```
Remembering that we don't want to install any database solution, the ideal would be to not even consider a container for that.
```

- A basic table structure should be created. The App_setting found must be added to this table allowing the user to be able to handle it later through the web interface.

---
**REF010** - This reference will be described later because it is something more complex.

---
**REF011** - This reference is related to REF010. A more suitable folder structure should be considered allowing us to build two different applications.

---
**REF012** - Builds must provide executables compatible with Linux, Windows and Mac OS.

---

## Team / Developers

- Rômulo Santos - [[@rromulos](https://github.com/rromulos)]

