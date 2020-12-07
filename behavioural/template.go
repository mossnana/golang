package behavioural

func TestTemplate() {
	// Send sms
	smsOTP := &sms{}
	oSMS := otp{
		iOTP: smsOTP,
	}
	oSMS.genAndSendOTP(4)
	// Send email
	emailOTP := &email{}
	oEmail := otp{
		iOTP: emailOTP,
	}
	oEmail.genAndSendOTP(4)
}

type iOTP interface {
	genRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
	publicMetric()
}

// OTP

type otp struct {
	iOTP
}

func (o *otp) genAndSendOTP(otpLength int) error {
	otp := o.iOTP.genRandomOTP(otpLength)
	o.iOTP.saveOTPCache(otp)
	message := o.iOTP.getMessage(otp)
	err := o.iOTP.sendNotification(message)
	if err != nil {
		return err
	}
	o.iOTP.publicMetric()
	return nil
}

// SMS

type sms struct {
	*otp
}

func (s *sms) genRandomOTP(len int) string {
	randomOTP := "1234"
	println("SMS : ", randomOTP)
	return randomOTP
}

func (s *sms) getMessage(otp string) string {
	return "SMS OTP is " + otp
}

func (s *sms) sendNotification(message string) error {
	println("SMS Sending : ", message)
	return nil
}

func (s *sms) publishMetric() {
	println("public metric")
}

// Email
type email struct {
	otp
}

func (s *email) genRandomOTP(len int) string {
	randomOTP := "4321"
	println("Email : ", randomOTP)
	return randomOTP
}

func (s *email) getMessage(otp string) string {
	return "Email OTP is " + otp
}

func (s *email) sendNotification(message string) error {
	println("Email Sending : ", message)
	return nil
}

func (s *email) publishMetric() {
	println("public metric")
}
