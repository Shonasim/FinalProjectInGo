package errors

import "errors"

var ErrInvalidFirstName = errors.New("empty or invalid first name")
var ErrInvalidLastName = errors.New("empty or invalid last name")
var ErrInvalidFathersName = errors.New("empty or invalid fathers name")
var ErrInvalidEmail = errors.New("empty or invalid email")
var ErrInvalidPassword = errors.New("empty or invalid password")
var ErrInvalidSex = errors.New("empty or invalid sex")
var ErrInvalidAboutUser = errors.New("empty or invalid info about User")
var ErrInvalidPhoto = errors.New("empty or invalid photo")
var ErrFailedHashing = errors.New("failed to hash")
var ErrRecordNotFound = errors.New("record not found")
var ErrInvalidModel = errors.New("empty or invalid Model")
var ErrInvalidMark = errors.New("empty or invalid mark")
var ErrInvalidAutobody = errors.New("empty or invalid autobody")
var ErrInvalidCarNumber = errors.New("empty or invalid car number")
var ErrInvalidSeats = errors.New("empty or invalid seats")
var ErrAlreadyExists = errors.New("user email already exists")
var ErrBindJSON = errors.New("failed to read data")
var ErrFailedCreate = errors.New("failed to create")
var ErrFailedCreateCar = errors.New("failed to create car")
var ErrGetCities = errors.New("failed to get cities")
