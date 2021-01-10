package random_test

import (
	"dataman/random"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRandom struct {
	mock.Mock
}

func (o *MockRandom) Seed(seed int64) {
	o.Called(seed)
}

func (o *MockRandom) Float64() float64 {
	args := o.Called()
	return args.Get(0).(float64)
}

func (o *MockRandom) Int63n(n int64) int64 {
	args := o.Called(n)
	return args.Get(0).(int64)
}

func TestCreatingRandomShouldInitialSeed(t *testing.T) {
	mockRandom := new(MockRandom)

	mockRandom.On("Seed", mock.AnythingOfType("int64"))

	random.New(mockRandom)

	mockRandom.AssertExpectations(t)
}

func TestCallingIntShouldCallInt63nWithGivenMax(t *testing.T) {
	mockMaxValue := int64(1000)
	expectedResultValue := int64(777)
	expectedMaxValue := int64(1000)

	mockRandom := new(MockRandom)

	mockRandom.On("Seed", mock.AnythingOfType("int64"))
	mockRandom.On("Int63n", expectedMaxValue).Return(expectedResultValue)

	r := random.New(mockRandom)
	result := r.Int(mockMaxValue)

	mockRandom.AssertExpectations(t)
	assert.Equal(t, expectedResultValue, result)
}

func TestCallingCloseIntShouldCallInt63nWithGivenMaxPlus1(t *testing.T) {
	mockMaxValue := int64(1000)
	expectedResultValue := int64(777)
	expectedMaxValue := int64(1001)

	mockRandom := new(MockRandom)

	mockRandom.On("Seed", mock.AnythingOfType("int64"))
	mockRandom.On("Int63n", expectedMaxValue).Return(expectedResultValue)

	r := random.New(mockRandom)
	result := r.CloseInt(mockMaxValue)

	mockRandom.AssertExpectations(t)
	assert.Equal(t, expectedResultValue, result)
}

func TestCallingIntBetweenShouldCallInt63nWithDiffBetweenMinAndMax(t *testing.T) {
	mockRandomValue := int64(777)
	expectedResultValue := int64(877)
	expectedMinValue := int64(100)
	expectedMaxValue := int64(1000)
	expectedDeltaValue := expectedMaxValue - expectedMinValue + 1

	mockRandom := new(MockRandom)

	mockRandom.On("Seed", mock.AnythingOfType("int64"))
	mockRandom.On("Int63n", expectedDeltaValue).Return(mockRandomValue)

	r := random.New(mockRandom)
	result := r.IntBetween(expectedMinValue, expectedMaxValue)

	mockRandom.AssertExpectations(t)
	assert.Equal(t, expectedResultValue, result)
}

func TestCallingIntBetweenWithInverseMinMaxShouldCallInt63nWithDiffBetweenMinAndMax(t *testing.T) {
	mockRandomValue := int64(777)
	expectedResultValue := int64(877)
	expectedMinValue := int64(1000)
	expectedMaxValue := int64(100)
	expectedDeltaValue := expectedMinValue - expectedMaxValue + 1

	mockRandom := new(MockRandom)

	mockRandom.On("Seed", mock.AnythingOfType("int64"))
	mockRandom.On("Int63n", expectedDeltaValue).Return(mockRandomValue)

	r := random.New(mockRandom)
	result := r.IntBetween(expectedMinValue, expectedMaxValue)

	mockRandom.AssertExpectations(t)
	assert.Equal(t, expectedResultValue, result)
}

func TestCallingIntBetweenWithMinAndMinPlus1ShouldNotCallInt63n(t *testing.T) {
	expectedResultValue := int64(100)
	expectedMinValue := int64(100)
	expectedMaxValue := int64(101)
	expectedDeltaValue := expectedMaxValue - expectedMinValue + 1

	mockRandom := new(MockRandom)

	mockRandom.On("Seed", mock.AnythingOfType("int64"))
	mockRandom.On("Int63n", expectedDeltaValue).Panic("Should have not been called!")

	r := random.New(mockRandom)
	result := r.IntBetween(expectedMinValue, expectedMaxValue)

	mockRandom.AssertNotCalled(t, "Int63n", 2)
	assert.Equal(t, expectedResultValue, result)
}

func TestCallingIntBetweenWithTheSameMinAndMaxShouldNotCallInt63n(t *testing.T) {
	expectedResultValue := int64(100)
	expectedMinValue := int64(100)
	expectedMaxValue := int64(100)
	expectedDeltaValue := expectedMaxValue - expectedMinValue + 1

	mockRandom := new(MockRandom)

	mockRandom.On("Seed", mock.AnythingOfType("int64"))
	mockRandom.On("Int63n", expectedDeltaValue).Panic("Should have not been called!")

	r := random.New(mockRandom)
	result := r.IntBetween(expectedMinValue, expectedMaxValue)

	mockRandom.AssertNotCalled(t, "Int63n", 1)
	assert.Equal(t, expectedResultValue, result)
}

func TestCallingCloseFloat(t *testing.T) {
	mockRandomValue1 := float64(0.1)
	mockRandomValue2 := float64(0.2)
	expectedResultValue := float64(0.5)

	mockRandom := new(MockRandom)

	mockRandom.On("Seed", mock.AnythingOfType("int64"))
	mockRandom.On("Float64").Return(mockRandomValue1).Once()
	mockRandom.On("Float64").Return(mockRandomValue2).Once()

	r := random.New(mockRandom)
	result := r.CloseFloat()

	mockRandom.AssertExpectations(t)
	assert.Equal(t, expectedResultValue, result)
}

func TestCallingCloseFloatWhenDivisorIsGreaterThanDividend(t *testing.T) {
	mockRandomValue1 := float64(0.2)
	mockRandomValue2 := float64(0.1)
	expectedResultValue := float64(0.5)

	mockRandom := new(MockRandom)

	mockRandom.On("Seed", mock.AnythingOfType("int64"))
	mockRandom.On("Float64").Return(mockRandomValue1).Once()
	mockRandom.On("Float64").Return(mockRandomValue2).Once()

	r := random.New(mockRandom)
	result := r.CloseFloat()

	mockRandom.AssertExpectations(t)
	assert.Equal(t, expectedResultValue, result)
}

func TestCallingCloseFloatWhenDivisorIsEqualToDividend(t *testing.T) {
	mockRandomValue1 := float64(0.1)
	mockRandomValue2 := float64(0.1)
	expectedResultValue := float64(1.0)

	mockRandom := new(MockRandom)

	mockRandom.On("Seed", mock.AnythingOfType("int64"))
	mockRandom.On("Float64").Return(mockRandomValue1).Once()
	mockRandom.On("Float64").Return(mockRandomValue2).Once()

	r := random.New(mockRandom)
	result := r.CloseFloat()

	mockRandom.AssertExpectations(t)
	assert.Equal(t, expectedResultValue, result)
}

func TestCallingCloseFloatWhenDivisorIsZero(t *testing.T) {
	mockRandomValue1 := float64(0.0)
	mockRandomValue2 := float64(0.0)
	expectedResultValue := float64(1.0)

	mockRandom := new(MockRandom)

	mockRandom.On("Seed", mock.AnythingOfType("int64"))
	mockRandom.On("Float64").Return(mockRandomValue1).Once()
	mockRandom.On("Float64").Return(mockRandomValue2).Once()

	r := random.New(mockRandom)
	result := r.CloseFloat()

	mockRandom.AssertExpectations(t)
	assert.Equal(t, expectedResultValue, result)
}

func TestCallingCloseFloatWhenDividendIsZero(t *testing.T) {
	mockRandomValue1 := float64(0.8)
	mockRandomValue2 := float64(0.0)
	expectedResultValue := float64(0.0)

	mockRandom := new(MockRandom)

	mockRandom.On("Seed", mock.AnythingOfType("int64"))
	mockRandom.On("Float64").Return(mockRandomValue1).Once()
	mockRandom.On("Float64").Return(mockRandomValue2).Once()

	r := random.New(mockRandom)
	result := r.CloseFloat()

	mockRandom.AssertExpectations(t)
	assert.Equal(t, expectedResultValue, result)
}

func TestCallingDecimalShouldReturnARandomProportionOfMax(t *testing.T) {
	mockMaxValue := float64(1000.0)
	mockRandomValue1 := float64(0.1)
	mockRandomValue2 := float64(0.2)
	expectedResultValue := float64(500.0)

	mockRandom := new(MockRandom)

	mockRandom.On("Seed", mock.AnythingOfType("int64"))
	mockRandom.On("Float64").Return(mockRandomValue1).Once()
	mockRandom.On("Float64").Return(mockRandomValue2).Once()

	r := random.New(mockRandom)
	result := r.Decimal(mockMaxValue)

	mockRandom.AssertExpectations(t)
	assert.Equal(t, expectedResultValue, result)
}

func TestCallingDecimalWithMinMaxShouldReturnARandomProportionOfDelta(t *testing.T) {
	mockMinValue := float64(100.0)
	mockMaxValue := float64(1000.0)
	mockRandomValue1 := float64(0.1)
	mockRandomValue2 := float64(0.2)
	expectedResultValue := float64(550.0)

	mockRandom := new(MockRandom)

	mockRandom.On("Seed", mock.AnythingOfType("int64"))
	mockRandom.On("Float64").Return(mockRandomValue1).Once()
	mockRandom.On("Float64").Return(mockRandomValue2).Once()

	r := random.New(mockRandom)
	result := r.DecimalBetween(mockMinValue, mockMaxValue)

	mockRandom.AssertExpectations(t)
	assert.Equal(t, expectedResultValue, result)
}

func TestCallingDecimalWithInverseMinMaxShouldReturnARandomProportionOfDelta(t *testing.T) {
	mockMinValue := float64(1000.0)
	mockMaxValue := float64(100.0)
	mockRandomValue1 := float64(0.1)
	mockRandomValue2 := float64(0.2)
	expectedResultValue := float64(550.0)

	mockRandom := new(MockRandom)

	mockRandom.On("Seed", mock.AnythingOfType("int64"))
	mockRandom.On("Float64").Return(mockRandomValue1).Once()
	mockRandom.On("Float64").Return(mockRandomValue2).Once()

	r := random.New(mockRandom)
	result := r.DecimalBetween(mockMinValue, mockMaxValue)

	mockRandom.AssertExpectations(t)
	assert.Equal(t, expectedResultValue, result)
}

func TestCallingElementShouldReturnTheElementAtTheReturnRandomValueIndex(t *testing.T) {
	mockArrayData := []string{
		"What",
		"A",
		"Wonderful",
		"World",
		".",
	}
	expectArraySize := int64(len(mockArrayData))
	mockRandomValue := int64(3)
	expectedResultValue := mockArrayData[3]

	mockRandom := new(MockRandom)

	mockRandom.On("Seed", mock.AnythingOfType("int64"))
	mockRandom.On("Int63n", expectArraySize).Return(mockRandomValue)

	r := random.New(mockRandom)
	result, e := r.Element(mockArrayData)

	mockRandom.AssertExpectations(t)
	assert.Nil(t, e)
	assert.Equal(t, expectedResultValue, result)
}

func TestCallingElementShouldReturnErrorIfGivenArgumentIsNotArray(t *testing.T) {
	mockParameter := 40
	mockRandomValue := int64(3)

	mockRandom := new(MockRandom)

	mockRandom.On("Seed", mock.AnythingOfType("int64"))
	mockRandom.On("Int63n", mock.AnythingOfType("int64")).Return(mockRandomValue)

	r := random.New(mockRandom)
	result, e := r.Element(mockParameter)

	mockRandom.AssertCalled(t, "Seed", mock.AnythingOfType("int64"))
	mockRandom.AssertNotCalled(t, "Int63n")
	assert.NotNil(t, e)
	assert.Nil(t, result)
}

func TestCallingElementShouldReturnErrorIfGivenArgumentIsEmptyArray(t *testing.T) {
	mockArrayData := []string{}
	expectArraySize := int64(len(mockArrayData))
	mockRandomValue := int64(3)

	mockRandom := new(MockRandom)

	mockRandom.On("Seed", mock.AnythingOfType("int64"))
	mockRandom.On("Int63n", expectArraySize).Return(mockRandomValue)

	r := random.New(mockRandom)
	result, e := r.Element(mockArrayData)

	mockRandom.AssertCalled(t, "Seed", mock.AnythingOfType("int64"))
	mockRandom.AssertNotCalled(t, "Int63n")
	assert.NotNil(t, e)
	assert.Nil(t, result)
}

func TestCallingDateShouldReturnTheRandomDateByRandomSecondSinceEpochWithMax99991231235959999999999UTC(t *testing.T) {
	expectedMaxSecValue := random.MaxRandomSec + 1
	expectedMaxNanosecValue := random.MaxRandomNanosec + 1
	mockRandomSecValue := int64(1555555555)
	mockRandomNanoValue := int64(2000000)
	expectedResultValue := time.Unix(mockRandomSecValue, mockRandomNanoValue)

	mockRandom := new(MockRandom)

	mockRandom.On("Seed", mock.AnythingOfType("int64"))
	mockRandom.On("Int63n", expectedMaxSecValue).Return(mockRandomSecValue).Once()
	mockRandom.On("Int63n", expectedMaxNanosecValue).Return(mockRandomNanoValue).Once()

	r := random.New(mockRandom)
	result := r.Date()

	mockRandom.AssertExpectations(t)
	assert.Equal(t, expectedResultValue, result)
}

func TestCallingDateBetweenShouldReturnTheRandomDateByRandomSecondOffsetBetweenDatesDelta(t *testing.T) {
	mockMinDate, _ := time.Parse(time.RFC3339Nano, "2006-01-02T00:00:00.000000000Z")
	mockMaxDate, _ := time.Parse(time.RFC3339Nano, "2007-12-31T23:59:59.999999999Z")
	expectedMaxNanosecValue := int64(62985600000000000)
	mockRandomNanoValue := int64(41234567890123456)
	expectedResultValue, _ := time.Parse(time.RFC3339Nano, "2007-04-24T06:02:47.890123456Z")

	mockRandom := new(MockRandom)

	mockRandom.On("Seed", mock.AnythingOfType("int64"))
	mockRandom.On("Int63n", expectedMaxNanosecValue).Return(mockRandomNanoValue)

	r := random.New(mockRandom)
	result := r.DateBetween(mockMinDate, mockMaxDate)

	mockRandom.AssertExpectations(t)
	assert.Equal(t, expectedResultValue, result)
}

func TestCallingDateBetweenWithInverseMinMaxShouldReturnTheRandomDateByRandomSecondOffsetBetweenDatesDelta(t *testing.T) {
	mockMinDate, _ := time.Parse(time.RFC3339Nano, "2007-12-31T23:59:59.999999999Z")
	mockMaxDate, _ := time.Parse(time.RFC3339Nano, "2006-01-02T00:00:00.000000000Z")
	expectedMaxNanosecValue := int64(62985600000000000)
	mockRandomNanoValue := int64(41234567890123456)
	expectedResultValue, _ := time.Parse(time.RFC3339Nano, "2007-04-24T06:02:47.890123456Z")

	mockRandom := new(MockRandom)

	mockRandom.On("Seed", mock.AnythingOfType("int64"))
	mockRandom.On("Int63n", expectedMaxNanosecValue).Return(mockRandomNanoValue)

	r := random.New(mockRandom)
	result := r.DateBetween(mockMinDate, mockMaxDate)

	mockRandom.AssertExpectations(t)
	assert.Equal(t, expectedResultValue, result)
}

func TestCallingDateRangeShouldReturnPeriod(t *testing.T) {
	expectedFromDate, _ := time.Parse(time.RFC3339Nano, "2006-04-28T00:46:40.651687623+07:00")
	expectedToDate, _ := time.Parse(time.RFC3339Nano, "2007-01-19T05:40:00.888177792+07:00")
	expectedResultValue := random.Period{
		From: expectedFromDate,
		To:   expectedToDate,
	}

	mockRandomFromSecValue := int64(1146160000)
	mockRandomFromNanoValue := int64(651687623)
	mockRandomToSecValue := int64(1169160000)
	mockRandomToNanoValue := int64(888177792)

	mockRandom := new(MockRandom)

	mockRandom.On("Seed", mock.AnythingOfType("int64"))
	mockRandom.On("Int63n", random.MaxRandomSec+1).Return(mockRandomFromSecValue).Once()
	mockRandom.On("Int63n", random.MaxRandomNanosec+1).Return(mockRandomFromNanoValue).Once()
	mockRandom.On("Int63n", random.MaxRandomSec+1).Return(mockRandomToSecValue).Once()
	mockRandom.On("Int63n", random.MaxRandomNanosec+1).Return(mockRandomToNanoValue).Once()

	r := random.New(mockRandom)
	result := r.DateRange()

	mockRandom.AssertExpectations(t)
	assert.Equal(t, expectedResultValue, result)
}

func TestCallingDateRangeShouldReturnPeriodCorrectlyEvenRandomFromDateIsAfterToDate(t *testing.T) {
	expectedFromDate, _ := time.Parse(time.RFC3339Nano, "2006-04-28T00:46:40.651687623+07:00")
	expectedToDate, _ := time.Parse(time.RFC3339Nano, "2007-01-19T05:40:00.888177792+07:00")
	expectedResultValue := random.Period{
		From: expectedFromDate,
		To:   expectedToDate,
	}

	mockRandomFromSecValue := int64(1146160000)
	mockRandomFromNanoValue := int64(651687623)
	mockRandomToSecValue := int64(1169160000)
	mockRandomToNanoValue := int64(888177792)

	mockRandom := new(MockRandom)

	mockRandom.On("Seed", mock.AnythingOfType("int64"))
	mockRandom.On("Int63n", random.MaxRandomSec+1).Return(mockRandomToSecValue).Once()
	mockRandom.On("Int63n", random.MaxRandomNanosec+1).Return(mockRandomToNanoValue).Once()
	mockRandom.On("Int63n", random.MaxRandomSec+1).Return(mockRandomFromSecValue).Once()
	mockRandom.On("Int63n", random.MaxRandomNanosec+1).Return(mockRandomFromNanoValue).Once()

	r := random.New(mockRandom)
	result := r.DateRange()

	mockRandom.AssertExpectations(t)
	assert.Equal(t, expectedResultValue, result)
}

func TestCallingDateRangeBetweenWithMinMaxShouldReturnTheRandomDatePeriod(t *testing.T) {
	mockMinDate, _ := time.Parse(time.RFC3339Nano, "2006-01-02T00:00:00.000000000Z")
	mockMaxDate, _ := time.Parse(time.RFC3339Nano, "2007-12-31T23:59:59.999999999Z")
	expectedMaxNanosecValue := int64(62985600000000000)
	mockRandomFromNanoValue := int64(41234567890123456)
	mockRandomToNanoValue := int64(42384567890123456)
	expectedFromDate, _ := time.Parse(time.RFC3339Nano, "2007-04-24T06:02:47.890123456Z")
	expectedToDate, _ := time.Parse(time.RFC3339Nano, "2007-05-07T13:29:27.890123456Z")
	expectedResultValue := random.Period{
		From: expectedFromDate,
		To:   expectedToDate,
	}

	mockRandom := new(MockRandom)

	mockRandom.On("Seed", mock.AnythingOfType("int64"))
	mockRandom.On("Int63n", expectedMaxNanosecValue).Return(mockRandomFromNanoValue).Once()
	mockRandom.On("Int63n", expectedMaxNanosecValue).Return(mockRandomToNanoValue).Once()

	r := random.New(mockRandom)
	result := r.DateRangeBetween(mockMinDate, mockMaxDate)

	mockRandom.AssertExpectations(t)
	assert.Equal(t, expectedResultValue, result)
}

func TestCallingDateRangeBetweenWithMinMaxShouldReturnTheRandomDatePeriodEvenRandomDeltaIsInverse(t *testing.T) {
	mockMinDate, _ := time.Parse(time.RFC3339Nano, "2006-01-02T00:00:00.000000000Z")
	mockMaxDate, _ := time.Parse(time.RFC3339Nano, "2007-12-31T23:59:59.999999999Z")
	expectedMaxNanosecValue := int64(62985600000000000)
	mockRandomFromNanoValue := int64(42384567890123456)
	mockRandomToNanoValue := int64(41234567890123456)
	expectedFromDate, _ := time.Parse(time.RFC3339Nano, "2007-04-24T06:02:47.890123456Z")
	expectedToDate, _ := time.Parse(time.RFC3339Nano, "2007-05-07T13:29:27.890123456Z")
	expectedResultValue := random.Period{
		From: expectedFromDate,
		To:   expectedToDate,
	}

	mockRandom := new(MockRandom)

	mockRandom.On("Seed", mock.AnythingOfType("int64"))
	mockRandom.On("Int63n", expectedMaxNanosecValue).Return(mockRandomFromNanoValue).Once()
	mockRandom.On("Int63n", expectedMaxNanosecValue).Return(mockRandomToNanoValue).Once()

	r := random.New(mockRandom)
	result := r.DateRangeBetween(mockMinDate, mockMaxDate)

	mockRandom.AssertExpectations(t)
	assert.Equal(t, expectedResultValue, result)
}
