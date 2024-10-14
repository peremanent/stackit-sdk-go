/*
STACKIT MongoDB Service API

This is the documentation for the STACKIT MongoDB Flex Service API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbflex

import (
	"encoding/json"
)

// checks if the BackupSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupSchedule{}

// BackupSchedule struct for BackupSchedule
type BackupSchedule struct {
	BackupSchedule                 *string `json:"backupSchedule,omitempty"`
	DailySnapshotRetentionDays     *int64  `json:"dailySnapshotRetentionDays,omitempty"`
	MonthlySnapshotRetentionMonths *int64  `json:"monthlySnapshotRetentionMonths,omitempty"`
	PointInTimeWindowHours         *int64  `json:"pointInTimeWindowHours,omitempty"`
	SnapshotRetentionDays          *int64  `json:"snapshotRetentionDays,omitempty"`
	WeeklySnapshotRetentionWeeks   *int64  `json:"weeklySnapshotRetentionWeeks,omitempty"`
}

// NewBackupSchedule instantiates a new BackupSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupSchedule() *BackupSchedule {
	this := BackupSchedule{}
	return &this
}

// NewBackupScheduleWithDefaults instantiates a new BackupSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupScheduleWithDefaults() *BackupSchedule {
	this := BackupSchedule{}
	return &this
}

// GetBackupSchedule returns the BackupSchedule field value if set, zero value otherwise.
func (o *BackupSchedule) GetBackupSchedule() *string {
	if o == nil || IsNil(o.BackupSchedule) {
		var ret *string
		return ret
	}
	return o.BackupSchedule
}

// GetBackupScheduleOk returns a tuple with the BackupSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSchedule) GetBackupScheduleOk() (*string, bool) {
	if o == nil || IsNil(o.BackupSchedule) {
		return nil, false
	}
	return o.BackupSchedule, true
}

// HasBackupSchedule returns a boolean if a field has been set.
func (o *BackupSchedule) HasBackupSchedule() bool {
	if o != nil && !IsNil(o.BackupSchedule) {
		return true
	}

	return false
}

// SetBackupSchedule gets a reference to the given string and assigns it to the BackupSchedule field.
func (o *BackupSchedule) SetBackupSchedule(v *string) {
	o.BackupSchedule = v
}

// GetDailySnapshotRetentionDays returns the DailySnapshotRetentionDays field value if set, zero value otherwise.
func (o *BackupSchedule) GetDailySnapshotRetentionDays() *int64 {
	if o == nil || IsNil(o.DailySnapshotRetentionDays) {
		var ret *int64
		return ret
	}
	return o.DailySnapshotRetentionDays
}

// GetDailySnapshotRetentionDaysOk returns a tuple with the DailySnapshotRetentionDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSchedule) GetDailySnapshotRetentionDaysOk() (*int64, bool) {
	if o == nil || IsNil(o.DailySnapshotRetentionDays) {
		return nil, false
	}
	return o.DailySnapshotRetentionDays, true
}

// HasDailySnapshotRetentionDays returns a boolean if a field has been set.
func (o *BackupSchedule) HasDailySnapshotRetentionDays() bool {
	if o != nil && !IsNil(o.DailySnapshotRetentionDays) {
		return true
	}

	return false
}

// SetDailySnapshotRetentionDays gets a reference to the given int64 and assigns it to the DailySnapshotRetentionDays field.
func (o *BackupSchedule) SetDailySnapshotRetentionDays(v *int64) {
	o.DailySnapshotRetentionDays = v
}

// GetMonthlySnapshotRetentionMonths returns the MonthlySnapshotRetentionMonths field value if set, zero value otherwise.
func (o *BackupSchedule) GetMonthlySnapshotRetentionMonths() *int64 {
	if o == nil || IsNil(o.MonthlySnapshotRetentionMonths) {
		var ret *int64
		return ret
	}
	return o.MonthlySnapshotRetentionMonths
}

// GetMonthlySnapshotRetentionMonthsOk returns a tuple with the MonthlySnapshotRetentionMonths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSchedule) GetMonthlySnapshotRetentionMonthsOk() (*int64, bool) {
	if o == nil || IsNil(o.MonthlySnapshotRetentionMonths) {
		return nil, false
	}
	return o.MonthlySnapshotRetentionMonths, true
}

// HasMonthlySnapshotRetentionMonths returns a boolean if a field has been set.
func (o *BackupSchedule) HasMonthlySnapshotRetentionMonths() bool {
	if o != nil && !IsNil(o.MonthlySnapshotRetentionMonths) {
		return true
	}

	return false
}

// SetMonthlySnapshotRetentionMonths gets a reference to the given int64 and assigns it to the MonthlySnapshotRetentionMonths field.
func (o *BackupSchedule) SetMonthlySnapshotRetentionMonths(v *int64) {
	o.MonthlySnapshotRetentionMonths = v
}

// GetPointInTimeWindowHours returns the PointInTimeWindowHours field value if set, zero value otherwise.
func (o *BackupSchedule) GetPointInTimeWindowHours() *int64 {
	if o == nil || IsNil(o.PointInTimeWindowHours) {
		var ret *int64
		return ret
	}
	return o.PointInTimeWindowHours
}

// GetPointInTimeWindowHoursOk returns a tuple with the PointInTimeWindowHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSchedule) GetPointInTimeWindowHoursOk() (*int64, bool) {
	if o == nil || IsNil(o.PointInTimeWindowHours) {
		return nil, false
	}
	return o.PointInTimeWindowHours, true
}

// HasPointInTimeWindowHours returns a boolean if a field has been set.
func (o *BackupSchedule) HasPointInTimeWindowHours() bool {
	if o != nil && !IsNil(o.PointInTimeWindowHours) {
		return true
	}

	return false
}

// SetPointInTimeWindowHours gets a reference to the given int64 and assigns it to the PointInTimeWindowHours field.
func (o *BackupSchedule) SetPointInTimeWindowHours(v *int64) {
	o.PointInTimeWindowHours = v
}

// GetSnapshotRetentionDays returns the SnapshotRetentionDays field value if set, zero value otherwise.
func (o *BackupSchedule) GetSnapshotRetentionDays() *int64 {
	if o == nil || IsNil(o.SnapshotRetentionDays) {
		var ret *int64
		return ret
	}
	return o.SnapshotRetentionDays
}

// GetSnapshotRetentionDaysOk returns a tuple with the SnapshotRetentionDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSchedule) GetSnapshotRetentionDaysOk() (*int64, bool) {
	if o == nil || IsNil(o.SnapshotRetentionDays) {
		return nil, false
	}
	return o.SnapshotRetentionDays, true
}

// HasSnapshotRetentionDays returns a boolean if a field has been set.
func (o *BackupSchedule) HasSnapshotRetentionDays() bool {
	if o != nil && !IsNil(o.SnapshotRetentionDays) {
		return true
	}

	return false
}

// SetSnapshotRetentionDays gets a reference to the given int64 and assigns it to the SnapshotRetentionDays field.
func (o *BackupSchedule) SetSnapshotRetentionDays(v *int64) {
	o.SnapshotRetentionDays = v
}

// GetWeeklySnapshotRetentionWeeks returns the WeeklySnapshotRetentionWeeks field value if set, zero value otherwise.
func (o *BackupSchedule) GetWeeklySnapshotRetentionWeeks() *int64 {
	if o == nil || IsNil(o.WeeklySnapshotRetentionWeeks) {
		var ret *int64
		return ret
	}
	return o.WeeklySnapshotRetentionWeeks
}

// GetWeeklySnapshotRetentionWeeksOk returns a tuple with the WeeklySnapshotRetentionWeeks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSchedule) GetWeeklySnapshotRetentionWeeksOk() (*int64, bool) {
	if o == nil || IsNil(o.WeeklySnapshotRetentionWeeks) {
		return nil, false
	}
	return o.WeeklySnapshotRetentionWeeks, true
}

// HasWeeklySnapshotRetentionWeeks returns a boolean if a field has been set.
func (o *BackupSchedule) HasWeeklySnapshotRetentionWeeks() bool {
	if o != nil && !IsNil(o.WeeklySnapshotRetentionWeeks) {
		return true
	}

	return false
}

// SetWeeklySnapshotRetentionWeeks gets a reference to the given int64 and assigns it to the WeeklySnapshotRetentionWeeks field.
func (o *BackupSchedule) SetWeeklySnapshotRetentionWeeks(v *int64) {
	o.WeeklySnapshotRetentionWeeks = v
}

func (o BackupSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BackupSchedule) {
		toSerialize["backupSchedule"] = o.BackupSchedule
	}
	if !IsNil(o.DailySnapshotRetentionDays) {
		toSerialize["dailySnapshotRetentionDays"] = o.DailySnapshotRetentionDays
	}
	if !IsNil(o.MonthlySnapshotRetentionMonths) {
		toSerialize["monthlySnapshotRetentionMonths"] = o.MonthlySnapshotRetentionMonths
	}
	if !IsNil(o.PointInTimeWindowHours) {
		toSerialize["pointInTimeWindowHours"] = o.PointInTimeWindowHours
	}
	if !IsNil(o.SnapshotRetentionDays) {
		toSerialize["snapshotRetentionDays"] = o.SnapshotRetentionDays
	}
	if !IsNil(o.WeeklySnapshotRetentionWeeks) {
		toSerialize["weeklySnapshotRetentionWeeks"] = o.WeeklySnapshotRetentionWeeks
	}
	return toSerialize, nil
}

type NullableBackupSchedule struct {
	value *BackupSchedule
	isSet bool
}

func (v NullableBackupSchedule) Get() *BackupSchedule {
	return v.value
}

func (v *NullableBackupSchedule) Set(val *BackupSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupSchedule(val *BackupSchedule) *NullableBackupSchedule {
	return &NullableBackupSchedule{value: val, isSet: true}
}

func (v NullableBackupSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
