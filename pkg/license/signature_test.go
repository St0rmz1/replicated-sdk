package license

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVerifySignature(t *testing.T) {
	tests := []struct {
		name       string
		license    string
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "basic valid signature",
			license: `apiVersion: kots.io/v1beta1
kind: License
metadata:
  name: testcustomer
spec:
  appSlug: my-app
  channelID: 1vusIYZLAVxMG6q760OJmRKj5i5
  channelName: My Channel
  customerName: Test Customer
  endpoint: https://replicated.app
  entitlements:
    bool_field:
      title: Bool Field
      value: true
      valueType: Boolean
    expires_at:
      description: License Expiration
      title: Expiration
      value: "2030-07-27T00:00:00Z"
      valueType: String
    hidden_field:
      isHidden: true
      title: Hidden Field
      value: this is secret
      valueType: String
    int_field:
      title: Int Field
      value: 123
      valueType: Integer
    string_field:
      title: StringField
      value: single line text
      valueType: String
    text_field:
      title: Text Field
      value: |-
        multi
        line
        text
      valueType: Text
  isAirgapSupported: true
  isGitOpsSupported: true
  isSnapshotSupported: true
  licenseID: 1vusOokxAVp1tkRGuyxnF23PJcq
  licenseSequence: 7
  licenseType: prod
  signature: eyJsaWNlbnNlRGF0YSI6ImV5SmhjR2xXWlhKemFXOXVJam9pYTI5MGN5NXBieTkyTVdKbGRHRXhJaXdpYTJsdVpDSTZJa3hwWTJWdWMyVWlMQ0p0WlhSaFpHRjBZU0k2ZXlKdVlXMWxJam9pZEdWemRHTjFjM1J2YldWeUluMHNJbk53WldNaU9uc2liR2xqWlc1elpVbEVJam9pTVhaMWMwOXZhM2hCVm5BeGRHdFNSM1Y1ZUc1R01qTlFTbU54SWl3aWJHbGpaVzV6WlZSNWNHVWlPaUp3Y205a0lpd2lZM1Z6ZEc5dFpYSk9ZVzFsSWpvaVZHVnpkQ0JEZFhOMGIyMWxjaUlzSW1Gd2NGTnNkV2NpT2lKdGVTMWhjSEFpTENKamFHRnVibVZzU1VRaU9pSXhkblZ6U1ZsYVRFRldlRTFITm5FM05qQlBTbTFTUzJvMWFUVWlMQ0pqYUdGdWJtVnNUbUZ0WlNJNklrMTVJRU5vWVc1dVpXd2lMQ0pzYVdObGJuTmxVMlZ4ZFdWdVkyVWlPamNzSW1WdVpIQnZhVzUwSWpvaWFIUjBjSE02THk5eVpYQnNhV05oZEdWa0xtRndjQ0lzSW1WdWRHbDBiR1Z0Wlc1MGN5STZleUppYjI5c1gyWnBaV3hrSWpwN0luUnBkR3hsSWpvaVFtOXZiQ0JHYVdWc1pDSXNJblpoYkhWbElqcDBjblZsTENKMllXeDFaVlI1Y0dVaU9pSkNiMjlzWldGdUluMHNJbVY0Y0dseVpYTmZZWFFpT25zaWRHbDBiR1VpT2lKRmVIQnBjbUYwYVc5dUlpd2laR1Z6WTNKcGNIUnBiMjRpT2lKTWFXTmxibk5sSUVWNGNHbHlZWFJwYjI0aUxDSjJZV3gxWlNJNklqSXdNekF0TURjdE1qZFVNREE2TURBNk1EQmFJaXdpZG1Gc2RXVlVlWEJsSWpvaVUzUnlhVzVuSW4wc0ltaHBaR1JsYmw5bWFXVnNaQ0k2ZXlKMGFYUnNaU0k2SWtocFpHUmxiaUJHYVdWc1pDSXNJblpoYkhWbElqb2lkR2hwY3lCcGN5QnpaV055WlhRaUxDSjJZV3gxWlZSNWNHVWlPaUpUZEhKcGJtY2lMQ0pwYzBocFpHUmxiaUk2ZEhKMVpYMHNJbWx1ZEY5bWFXVnNaQ0k2ZXlKMGFYUnNaU0k2SWtsdWRDQkdhV1ZzWkNJc0luWmhiSFZsSWpveE1qTXNJblpoYkhWbFZIbHdaU0k2SWtsdWRHVm5aWElpZlN3aWMzUnlhVzVuWDJacFpXeGtJanA3SW5ScGRHeGxJam9pVTNSeWFXNW5SbWxsYkdRaUxDSjJZV3gxWlNJNkluTnBibWRzWlNCc2FXNWxJSFJsZUhRaUxDSjJZV3gxWlZSNWNHVWlPaUpUZEhKcGJtY2lmU3dpZEdWNGRGOW1hV1ZzWkNJNmV5SjBhWFJzWlNJNklsUmxlSFFnUm1sbGJHUWlMQ0oyWVd4MVpTSTZJbTExYkhScFhHNXNhVzVsWEc1MFpYaDBJaXdpZG1Gc2RXVlVlWEJsSWpvaVZHVjRkQ0o5ZlN3aWFYTkJhWEpuWVhCVGRYQndiM0owWldRaU9uUnlkV1VzSW1selIybDBUM0J6VTNWd2NHOXlkR1ZrSWpwMGNuVmxMQ0pwYzFOdVlYQnphRzkwVTNWd2NHOXlkR1ZrSWpwMGNuVmxmWDA9IiwiaW5uZXJTaWduYXR1cmUiOiJleUpzYVdObGJuTmxVMmxuYm1GMGRYSmxJam9pYUhneE1XTXZUR1ozUTNoVE5YRmtRWEJGU1hGdVRrMU9NMHBLYTJzNFZHZFhSVVpzVDFKVlJ6UjJjR1YzZEZoV1YzbG1lamRZY0hBd1ExazJZamRyUVRSS2N6TklhR3d3YkZJMFdUQTFMemN2UVVkQ2FEZFZNSGczUkhaTVozUXpVM00wYm5GTFZTdFhXRXBTVHpKWVFVRnZSME4xZFRWR1RGcHJRVWhYY1RSUVFtMXphSFY2Y1ZsdmNucHhlbGhGWVZWVlpFUlVkVXhDTW1nNWFIZ3dXRWhQUmxwUk16bHVkbTlPUjJaT2R5OTRTVmRaZEhSUGRYZHZhMncyTVZsb1JVeFZlRmQxU1ZSRmMwTlVhM2xtTVRNd09IazVSbFJzWlRKeVYyZEVlSEZNYTBSUFNXVXlPRWwzUzJSQkwySXdWVUl5VEZGbVRWcHdWemwyUTNCSkwybHlWek5uYmpaeU5WWjNWMjB2U1dweWJtNDNSelJrVmpadVYzcFRkMGhQUTJSdWEwMTRNRXQ1VVVOa0wxQjFaWEpUYjNSdVEwOXRTMDEzWlRSTGJqaERkMU5YVVRRNGRURkRNbTFpV1VzeGRYTlpOM1YzUFQwaUxDSndkV0pzYVdOTFpYa2lPaUl0TFMwdExVSkZSMGxPSUZCVlFreEpReUJMUlZrdExTMHRMVnh1VFVsSlFrbHFRVTVDWjJ0eGFHdHBSemwzTUVKQlVVVkdRVUZQUTBGUk9FRk5TVWxDUTJkTFEwRlJSVUZ6TkhKdlVIcDFhV1JNZVhOMmIxWTJkemxhTkZ4dVdHRmliME5tWTJNeGFHZFZhQ3N3V1VkS2NFNURSVXhyTjBaTFF5OTJhemR6ZERsR05tY3dUMjlrU0VSbGVYZFJXa2hLZFU1TVpsUnNRbEJHUTJOaU5seHVObTlzVEZOeWNGQTRjbFUzU0d4SGJsRkVSMFJNYVhkS1EyaGtSRGRVVUdSM2FXdHBkMHRGY201aldqaEdaalZsU25vd2RETmlUWFpyVDJaVVluSkJiRnh1WWtGQ1kwbzVNVmxVT1hKdVVXOXFkVWN4UldKUVRqaEZWblI2TWxZNE5IZHViR2Q0TUhCd2JEVjRPSFpOYlhwcE1ISnVibEZVV1VGamJ6WnFhMnBJTTF4dVRuTlVkWE4xUzFkdlJGUjVNWE5yZGtSUk9IbEJZV0ptWTNNME4zWnNRazAwU0RGT1JFNHZSSFJhWWxZdllubDJia0o2YkM4eFZrVnpURmRqWlZWcFRGeHVSWEYxT0VkeWF5dFFVRGQyUkdSd2JFUjNjWFpQV2t4RmRYazNkamhuUm01U09WUlVSV3ByTlVvNWRuWlVTR2RtU25VemVubEVPR2xLWTBSRE5YcHFPVnh1YjFGSlJFRlJRVUpjYmkwdExTMHRSVTVFSUZCVlFreEpReUJMUlZrdExTMHRMVnh1SWl3aWEyVjVVMmxuYm1GMGRYSmxJam9pWlhsS2VtRlhaSFZaV0ZJeFkyMVZhVTlwU2pCUldIQjJXVE5LVms1NmFGaFNSMlJzVVRKb2NtTklXa1ZVVlRsRldqQktXVTFGUmtaVFJFNUZVMGhLYkUxclRUTkxNSEJFVkROR2VGTnROVVJVVlRWVlltMDFiVnBGUm5sWldIQjZaRVJqTVZaSGFFeFBXRUpVVWtacmRrd3diek5aTUZaSlVteFdWRXd5T1VoV1JXeHNWa1ZPTUZSSE1WWlJNR04zVkd4R2JGa3pTblJUUm1zMFZVWk9hMVpWU2pCVU1WbDNZbXQwY0ZSclZuQmpia0poVFZjNWFtSldiSEZaYTNob1UyeHNWV0pGUmtWWGJVWnZWakZLVUZkcWJGSmhXRVp1V2xkb1EyRnVRak5TUjNNd1lWWkpOVTVXVmxkV1ZUVnlUMGhLYjFsVlRYbGhiVGcwVjBkYWVGbHFWbFppYlhoeFpFWkZkMDU1Y3pCaFZsSkpWRVpPTm1WRk1IcGxWWFJ2VFVaR1ZtRXdWVFJSVnpsSFVsaEtVRTFZUmxCU01WcFJVMVJDTmxsV2FIcFdWWEJ0WTBSU2JFMVVRazlPVjNSU1ZucFdUMU5XWTNaU1ZYUkZVMGhzYlU5VmJGaGtNMUl3WTFWc1lXTlhSakJTYTA1RVlVWmtjbUo2VmtSU00wSllUREkxUmsxWVl6SmxWM1JKVlZoQk1sVXhTbEppU0Zwd1VrVXdNRlpFVWt0VU1rWnNVVmQwYzFSV1VrMVVWV055V1RCYVRHSXpaRTlUVm05NVlraE9SR1JzVG5aUmFrWmFaVmRPVGxOVlNteGFiRXB1Wld0U2RVMHhSVGxRVTBselNXMWtjMkl5U21oaVJYUnNaVlZzYTBscWIybFpiVkpzV2xSVk1rNVVXWGRaTWxwcFRrUk9hazlYU1hsUFIwcHRUMVJvYkZsWFRtaGFiVVV5VGtSWmFXWlJQVDBpZlE9PSJ9
`,
			wantErr:    false,
			wantErrMsg: "",
		},
		{
			name: "invalid signature",
			license: `apiVersion: kots.io/v1beta1
kind: License
metadata:
  name: testcustomer
spec:
  appSlug: my-app
  channelID: 1vusIYZLAVxMG6q760OJmRKj5i5
  channelName: My Channel
  customerName: Test Customer
  endpoint: https://replicated.app
  entitlements:
    bool_field:
      title: Bool Field
      value: true
      valueType: Boolean
    expires_at:
      description: License Expiration
      title: Expiration
      value: "2030-07-27T00:00:00Z"
      valueType: String
    hidden_field:
      isHidden: true
      title: Hidden Field
      value: this is secret
      valueType: String
    int_field:
      title: Int Field
      value: 123
      valueType: Integer
    string_field:
      title: StringField
      value: single line text
      valueType: String
    text_field:
      title: Text Field
      value: |-
        multi
        line
        text
      valueType: Text
  isAirgapSupported: true
  isGitOpsSupported: true
  isSnapshotSupported: true
  licenseID: 1vusOokxAVp1tkRGuyxnF23PJcq
  licenseSequence: 7
  licenseType: prod
  signature: eyJsaWNlbnNlRGF0YSI6ImV5SmhjR2xXWlhKemFXOXVJam9pYTI5MGN5NXBieTkyTVdKbGRHRXhJaXdpYTJsdVpDSTZJa3hwWTJWdWMyVWlMQ0p0WlhSaFpHRjBZU0k2ZXlKdVlXMWxJam9pZEdWemRHTjFjM1J2YldWeUluMHNJbk53WldNaU9uc2liR2xqWlc1elpVbEVJam9pTVhaMWMwOXZhM2hCVm5BeGRHdFNSM1Y1ZUc1R01qTlFTbU54SWl3aWJHbGpaVzV6WlZSNWNHVWlPaUp3Y205a0lpd2lZM1Z6ZEc5dFpYSk9ZVzFsSWpvaVZHVnpkQ0JEZFhOMGIyMWxjaUlzSW1Gd2NGTnNkV2NpT2lKdGVTMWhjSEF0Ylc5a2FXWnBaV1FpTENKamFHRnVibVZzU1VRaU9pSXhkblZ6U1ZsYVRFRldlRTFITm5FM05qQlBTbTFTUzJvMWFUVWlMQ0pqYUdGdWJtVnNUbUZ0WlNJNklrMTVJRU5vWVc1dVpXd2lMQ0pzYVdObGJuTmxVMlZ4ZFdWdVkyVWlPamNzSW1WdVpIQnZhVzUwSWpvaWFIUjBjSE02THk5eVpYQnNhV05oZEdWa0xtRndjQzV0YjJScFptbGxaQ0lzSW1WdWRHbDBiR1Z0Wlc1MGN5STZleUppYjI5c1gyWnBaV3hrSWpwN0luUnBkR3hsSWpvaVFtOXZiQ0JHYVdWc1pDSXNJblpoYkhWbElqcDBjblZsTENKMllXeDFaVlI1Y0dVaU9pSkNiMjlzWldGdUluMHNJbVY0Y0dseVpYTmZZWFFpT25zaWRHbDBiR1VpT2lKRmVIQnBjbUYwYVc5dUlpd2laR1Z6WTNKcGNIUnBiMjRpT2lKTWFXTmxibk5sSUVWNGNHbHlZWFJwYjI0aUxDSjJZV3gxWlNJNklqSXdNekF0TURjdE1qZFVNREE2TURBNk1EQmFJaXdpZG1Gc2RXVlVlWEJsSWpvaVUzUnlhVzVuSW4wc0ltaHBaR1JsYmw5bWFXVnNaQ0k2ZXlKMGFYUnNaU0k2SWtocFpHUmxiaUJHYVdWc1pDSXNJblpoYkhWbElqb2lkR2hwY3lCcGN5QnpaV055WlhRaUxDSjJZV3gxWlZSNWNHVWlPaUpUZEhKcGJtY2lMQ0pwYzBocFpHUmxiaUk2ZEhKMVpYMHNJbWx1ZEY5bWFXVnNaQ0k2ZXlKMGFYUnNaU0k2SWtsdWRDQkdhV1ZzWkNJc0luWmhiSFZsSWpveE1qTXNJblpoYkhWbFZIbHdaU0k2SWtsdWRHVm5aWElpZlN3aWMzUnlhVzVuWDJacFpXeGtJanA3SW5ScGRHeGxJam9pVTNSeWFXNW5SbWxsYkdRaUxDSjJZV3gxWlNJNkluTnBibWRzWlNCc2FXNWxJSFJsZUhRaUxDSjJZV3gxWlZSNWNHVWlPaUpUZEhKcGJtY2lmU3dpZEdWNGRGOW1hV1ZzWkNJNmV5SjBhWFJzWlNJNklsUmxlSFFnUm1sbGJHUWlMQ0oyWVd4MVpTSTZJbTExYkhScFhHNXNhVzVsWEc1MFpYaDBJaXdpZG1Gc2RXVlVlWEJsSWpvaVZHVjRkQ0o5ZlN3aWFYTkJhWEpuWVhCVGRYQndiM0owWldRaU9uUnlkV1VzSW1selIybDBUM0J6VTNWd2NHOXlkR1ZrSWpwMGNuVmxMQ0pwYzFOdVlYQnphRzkwVTNWd2NHOXlkR1ZrSWpwMGNuVmxmWDA9IiwiaW5uZXJTaWduYXR1cmUiOiJleUpzYVdObGJuTmxVMmxuYm1GMGRYSmxJam9pYUhneE1XTXZUR1ozUTNoVE5YRmtRWEJGU1hGdVRrMU9NMHBLYTJzNFZHZFhSVVpzVDFKVlJ6UjJjR1YzZEZoV1YzbG1lamRZY0hBd1ExazJZamRyUVRSS2N6TklhR3d3YkZJMFdUQTFMemN2UVVkQ2FEZFZNSGczUkhaTVozUXpVM00wYm5GTFZTdFhXRXBTVHpKWVFVRnZSME4xZFRWR1RGcHJRVWhYY1RSUVFtMXphSFY2Y1ZsdmNucHhlbGhGWVZWVlpFUlVkVXhDTW1nNWFIZ3dXRWhQUmxwUk16bHVkbTlPUjJaT2R5OTRTVmRaZEhSUGRYZHZhMncyTVZsb1JVeFZlRmQxU1ZSRmMwTlVhM2xtTVRNd09IazVSbFJzWlRKeVYyZEVlSEZNYTBSUFNXVXlPRWwzUzJSQkwySXdWVUl5VEZGbVRWcHdWemwyUTNCSkwybHlWek5uYmpaeU5WWjNWMjB2U1dweWJtNDNSelJrVmpadVYzcFRkMGhQUTJSdWEwMTRNRXQ1VVVOa0wxQjFaWEpUYjNSdVEwOXRTMDEzWlRSTGJqaERkMU5YVVRRNGRURkRNbTFpV1VzeGRYTlpOM1YzUFQwaUxDSndkV0pzYVdOTFpYa2lPaUl0TFMwdExVSkZSMGxPSUZCVlFreEpReUJMUlZrdExTMHRMVnh1VFVsSlFrbHFRVTVDWjJ0eGFHdHBSemwzTUVKQlVVVkdRVUZQUTBGUk9FRk5TVWxDUTJkTFEwRlJSVUZ6TkhKdlVIcDFhV1JNZVhOMmIxWTJkemxhTkZ4dVdHRmliME5tWTJNeGFHZFZhQ3N3V1VkS2NFNURSVXhyTjBaTFF5OTJhemR6ZERsR05tY3dUMjlrU0VSbGVYZFJXa2hLZFU1TVpsUnNRbEJHUTJOaU5seHVObTlzVEZOeWNGQTRjbFUzU0d4SGJsRkVSMFJNYVhkS1EyaGtSRGRVVUdSM2FXdHBkMHRGY201aldqaEdaalZsU25vd2RETmlUWFpyVDJaVVluSkJiRnh1WWtGQ1kwbzVNVmxVT1hKdVVXOXFkVWN4UldKUVRqaEZWblI2TWxZNE5IZHViR2Q0TUhCd2JEVjRPSFpOYlhwcE1ISnVibEZVV1VGamJ6WnFhMnBJTTF4dVRuTlVkWE4xUzFkdlJGUjVNWE5yZGtSUk9IbEJZV0ptWTNNME4zWnNRazAwU0RGT1JFNHZSSFJhWWxZdllubDJia0o2YkM4eFZrVnpURmRqWlZWcFRGeHVSWEYxT0VkeWF5dFFVRGQyUkdSd2JFUjNjWFpQV2t4RmRYazNkamhuUm01U09WUlVSV3ByTlVvNWRuWlVTR2RtU25VemVubEVPR2xLWTBSRE5YcHFPVnh1YjFGSlJFRlJRVUpjYmkwdExTMHRSVTVFSUZCVlFreEpReUJMUlZrdExTMHRMVnh1SWl3aWEyVjVVMmxuYm1GMGRYSmxJam9pWlhsS2VtRlhaSFZaV0ZJeFkyMVZhVTlwU2pCUldIQjJXVE5LVms1NmFGaFNSMlJzVVRKb2NtTklXa1ZVVlRsRldqQktXVTFGUmtaVFJFNUZVMGhLYkUxclRUTkxNSEJFVkROR2VGTnROVVJVVlRWVlltMDFiVnBGUm5sWldIQjZaRVJqTVZaSGFFeFBXRUpVVWtacmRrd3diek5aTUZaSlVteFdWRXd5T1VoV1JXeHNWa1ZPTUZSSE1WWlJNR04zVkd4R2JGa3pTblJUUm1zMFZVWk9hMVpWU2pCVU1WbDNZbXQwY0ZSclZuQmpia0poVFZjNWFtSldiSEZaYTNob1UyeHNWV0pGUmtWWGJVWnZWakZLVUZkcWJGSmhXRVp1V2xkb1EyRnVRak5TUjNNd1lWWkpOVTVXVmxkV1ZUVnlUMGhLYjFsVlRYbGhiVGcwVjBkYWVGbHFWbFppYlhoeFpFWkZkMDU1Y3pCaFZsSkpWRVpPTm1WRk1IcGxWWFJ2VFVaR1ZtRXdWVFJSVnpsSFVsaEtVRTFZUmxCU01WcFJVMVJDTmxsV2FIcFdWWEJ0WTBSU2JFMVVRazlPVjNSU1ZucFdUMU5XWTNaU1ZYUkZVMGhzYlU5VmJGaGtNMUl3WTFWc1lXTlhSakJTYTA1RVlVWmtjbUo2VmtSU00wSllUREkxUmsxWVl6SmxWM1JKVlZoQk1sVXhTbEppU0Zwd1VrVXdNRlpFVWt0VU1rWnNVVmQwYzFSV1VrMVVWV055V1RCYVRHSXpaRTlUVm05NVlraE9SR1JzVG5aUmFrWmFaVmRPVGxOVlNteGFiRXB1Wld0U2RVMHhSVGxRVTBselNXMWtjMkl5U21oaVJYUnNaVlZzYTBscWIybFpiVkpzV2xSVk1rNVVXWGRaTWxwcFRrUk9hazlYU1hsUFIwcHRUMVJvYkZsWFRtaGFiVVV5VGtSWmFXWlJQVDBpZlE9PSJ9
`,
			wantErr:    true,
			wantErrMsg: "failed to verify license signature: crypto/rsa: verification error: signature is invalid",
		},
		{
			name: "licenseID field changed",
			license: `apiVersion: kots.io/v1beta1
kind: License
metadata:
  name: testcustomer
spec:
  appSlug: my-app
  channelID: 1vusIYZLAVxMG6q760OJmRKj5i5
  channelName: My Channel
  customerName: Test Customer
  endpoint: https://replicated.app
  entitlements:
    bool_field:
      title: Bool Field
      value: true
      valueType: Boolean
    expires_at:
      description: License Expiration
      title: Expiration
      value: "2030-07-27T00:00:00Z"
      valueType: String
    hidden_field:
      isHidden: true
      title: Hidden Field
      value: this is secret
      valueType: String
    int_field:
      title: Int Field
      value: 123
      valueType: Integer
    string_field:
      title: StringField
      value: single line text
      valueType: String
    text_field:
      title: Text Field
      value: |-
        multi
        line
        text
      valueType: Text
  isAirgapSupported: true
  isGitOpsSupported: true
  isSnapshotSupported: true
  licenseID: 1vusOokxAVp1tkRGuyxnF23PJcq-modified
  licenseSequence: 7
  licenseType: prod
  signature: eyJsaWNlbnNlRGF0YSI6ImV5SmhjR2xXWlhKemFXOXVJam9pYTI5MGN5NXBieTkyTVdKbGRHRXhJaXdpYTJsdVpDSTZJa3hwWTJWdWMyVWlMQ0p0WlhSaFpHRjBZU0k2ZXlKdVlXMWxJam9pZEdWemRHTjFjM1J2YldWeUluMHNJbk53WldNaU9uc2liR2xqWlc1elpVbEVJam9pTVhaMWMwOXZhM2hCVm5BeGRHdFNSM1Y1ZUc1R01qTlFTbU54SWl3aWJHbGpaVzV6WlZSNWNHVWlPaUp3Y205a0lpd2lZM1Z6ZEc5dFpYSk9ZVzFsSWpvaVZHVnpkQ0JEZFhOMGIyMWxjaUlzSW1Gd2NGTnNkV2NpT2lKdGVTMWhjSEFpTENKamFHRnVibVZzU1VRaU9pSXhkblZ6U1ZsYVRFRldlRTFITm5FM05qQlBTbTFTUzJvMWFUVWlMQ0pqYUdGdWJtVnNUbUZ0WlNJNklrMTVJRU5vWVc1dVpXd2lMQ0pzYVdObGJuTmxVMlZ4ZFdWdVkyVWlPamNzSW1WdVpIQnZhVzUwSWpvaWFIUjBjSE02THk5eVpYQnNhV05oZEdWa0xtRndjQ0lzSW1WdWRHbDBiR1Z0Wlc1MGN5STZleUppYjI5c1gyWnBaV3hrSWpwN0luUnBkR3hsSWpvaVFtOXZiQ0JHYVdWc1pDSXNJblpoYkhWbElqcDBjblZsTENKMllXeDFaVlI1Y0dVaU9pSkNiMjlzWldGdUluMHNJbVY0Y0dseVpYTmZZWFFpT25zaWRHbDBiR1VpT2lKRmVIQnBjbUYwYVc5dUlpd2laR1Z6WTNKcGNIUnBiMjRpT2lKTWFXTmxibk5sSUVWNGNHbHlZWFJwYjI0aUxDSjJZV3gxWlNJNklqSXdNekF0TURjdE1qZFVNREE2TURBNk1EQmFJaXdpZG1Gc2RXVlVlWEJsSWpvaVUzUnlhVzVuSW4wc0ltaHBaR1JsYmw5bWFXVnNaQ0k2ZXlKMGFYUnNaU0k2SWtocFpHUmxiaUJHYVdWc1pDSXNJblpoYkhWbElqb2lkR2hwY3lCcGN5QnpaV055WlhRaUxDSjJZV3gxWlZSNWNHVWlPaUpUZEhKcGJtY2lMQ0pwYzBocFpHUmxiaUk2ZEhKMVpYMHNJbWx1ZEY5bWFXVnNaQ0k2ZXlKMGFYUnNaU0k2SWtsdWRDQkdhV1ZzWkNJc0luWmhiSFZsSWpveE1qTXNJblpoYkhWbFZIbHdaU0k2SWtsdWRHVm5aWElpZlN3aWMzUnlhVzVuWDJacFpXeGtJanA3SW5ScGRHeGxJam9pVTNSeWFXNW5SbWxsYkdRaUxDSjJZV3gxWlNJNkluTnBibWRzWlNCc2FXNWxJSFJsZUhRaUxDSjJZV3gxWlZSNWNHVWlPaUpUZEhKcGJtY2lmU3dpZEdWNGRGOW1hV1ZzWkNJNmV5SjBhWFJzWlNJNklsUmxlSFFnUm1sbGJHUWlMQ0oyWVd4MVpTSTZJbTExYkhScFhHNXNhVzVsWEc1MFpYaDBJaXdpZG1Gc2RXVlVlWEJsSWpvaVZHVjRkQ0o5ZlN3aWFYTkJhWEpuWVhCVGRYQndiM0owWldRaU9uUnlkV1VzSW1selIybDBUM0J6VTNWd2NHOXlkR1ZrSWpwMGNuVmxMQ0pwYzFOdVlYQnphRzkwVTNWd2NHOXlkR1ZrSWpwMGNuVmxmWDA9IiwiaW5uZXJTaWduYXR1cmUiOiJleUpzYVdObGJuTmxVMmxuYm1GMGRYSmxJam9pYUhneE1XTXZUR1ozUTNoVE5YRmtRWEJGU1hGdVRrMU9NMHBLYTJzNFZHZFhSVVpzVDFKVlJ6UjJjR1YzZEZoV1YzbG1lamRZY0hBd1ExazJZamRyUVRSS2N6TklhR3d3YkZJMFdUQTFMemN2UVVkQ2FEZFZNSGczUkhaTVozUXpVM00wYm5GTFZTdFhXRXBTVHpKWVFVRnZSME4xZFRWR1RGcHJRVWhYY1RSUVFtMXphSFY2Y1ZsdmNucHhlbGhGWVZWVlpFUlVkVXhDTW1nNWFIZ3dXRWhQUmxwUk16bHVkbTlPUjJaT2R5OTRTVmRaZEhSUGRYZHZhMncyTVZsb1JVeFZlRmQxU1ZSRmMwTlVhM2xtTVRNd09IazVSbFJzWlRKeVYyZEVlSEZNYTBSUFNXVXlPRWwzUzJSQkwySXdWVUl5VEZGbVRWcHdWemwyUTNCSkwybHlWek5uYmpaeU5WWjNWMjB2U1dweWJtNDNSelJrVmpadVYzcFRkMGhQUTJSdWEwMTRNRXQ1VVVOa0wxQjFaWEpUYjNSdVEwOXRTMDEzWlRSTGJqaERkMU5YVVRRNGRURkRNbTFpV1VzeGRYTlpOM1YzUFQwaUxDSndkV0pzYVdOTFpYa2lPaUl0TFMwdExVSkZSMGxPSUZCVlFreEpReUJMUlZrdExTMHRMVnh1VFVsSlFrbHFRVTVDWjJ0eGFHdHBSemwzTUVKQlVVVkdRVUZQUTBGUk9FRk5TVWxDUTJkTFEwRlJSVUZ6TkhKdlVIcDFhV1JNZVhOMmIxWTJkemxhTkZ4dVdHRmliME5tWTJNeGFHZFZhQ3N3V1VkS2NFNURSVXhyTjBaTFF5OTJhemR6ZERsR05tY3dUMjlrU0VSbGVYZFJXa2hLZFU1TVpsUnNRbEJHUTJOaU5seHVObTlzVEZOeWNGQTRjbFUzU0d4SGJsRkVSMFJNYVhkS1EyaGtSRGRVVUdSM2FXdHBkMHRGY201aldqaEdaalZsU25vd2RETmlUWFpyVDJaVVluSkJiRnh1WWtGQ1kwbzVNVmxVT1hKdVVXOXFkVWN4UldKUVRqaEZWblI2TWxZNE5IZHViR2Q0TUhCd2JEVjRPSFpOYlhwcE1ISnVibEZVV1VGamJ6WnFhMnBJTTF4dVRuTlVkWE4xUzFkdlJGUjVNWE5yZGtSUk9IbEJZV0ptWTNNME4zWnNRazAwU0RGT1JFNHZSSFJhWWxZdllubDJia0o2YkM4eFZrVnpURmRqWlZWcFRGeHVSWEYxT0VkeWF5dFFVRGQyUkdSd2JFUjNjWFpQV2t4RmRYazNkamhuUm01U09WUlVSV3ByTlVvNWRuWlVTR2RtU25VemVubEVPR2xLWTBSRE5YcHFPVnh1YjFGSlJFRlJRVUpjYmkwdExTMHRSVTVFSUZCVlFreEpReUJMUlZrdExTMHRMVnh1SWl3aWEyVjVVMmxuYm1GMGRYSmxJam9pWlhsS2VtRlhaSFZaV0ZJeFkyMVZhVTlwU2pCUldIQjJXVE5LVms1NmFGaFNSMlJzVVRKb2NtTklXa1ZVVlRsRldqQktXVTFGUmtaVFJFNUZVMGhLYkUxclRUTkxNSEJFVkROR2VGTnROVVJVVlRWVlltMDFiVnBGUm5sWldIQjZaRVJqTVZaSGFFeFBXRUpVVWtacmRrd3diek5aTUZaSlVteFdWRXd5T1VoV1JXeHNWa1ZPTUZSSE1WWlJNR04zVkd4R2JGa3pTblJUUm1zMFZVWk9hMVpWU2pCVU1WbDNZbXQwY0ZSclZuQmpia0poVFZjNWFtSldiSEZaYTNob1UyeHNWV0pGUmtWWGJVWnZWakZLVUZkcWJGSmhXRVp1V2xkb1EyRnVRak5TUjNNd1lWWkpOVTVXVmxkV1ZUVnlUMGhLYjFsVlRYbGhiVGcwVjBkYWVGbHFWbFppYlhoeFpFWkZkMDU1Y3pCaFZsSkpWRVpPTm1WRk1IcGxWWFJ2VFVaR1ZtRXdWVFJSVnpsSFVsaEtVRTFZUmxCU01WcFJVMVJDTmxsV2FIcFdWWEJ0WTBSU2JFMVVRazlPVjNSU1ZucFdUMU5XWTNaU1ZYUkZVMGhzYlU5VmJGaGtNMUl3WTFWc1lXTlhSakJTYTA1RVlVWmtjbUo2VmtSU00wSllUREkxUmsxWVl6SmxWM1JKVlZoQk1sVXhTbEppU0Zwd1VrVXdNRlpFVWt0VU1rWnNVVmQwYzFSV1VrMVVWV055V1RCYVRHSXpaRTlUVm05NVlraE9SR1JzVG5aUmFrWmFaVmRPVGxOVlNteGFiRXB1Wld0U2RVMHhSVGxRVTBselNXMWtjMkl5U21oaVJYUnNaVlZzYTBscWIybFpiVkpzV2xSVk1rNVVXWGRaTWxwcFRrUk9hazlYU1hsUFIwcHRUMVJvYkZsWFRtaGFiVVV5VGtSWmFXWlJQVDBpZlE9PSJ9
`,
			wantErr:    true,
			wantErrMsg: `"licenseID" field has changed`,
		},
		{
			name: "endpoint field changed",
			license: `apiVersion: kots.io/v1beta1
kind: License
metadata:
  name: testcustomer
spec:
  appSlug: my-app
  channelID: 1vusIYZLAVxMG6q760OJmRKj5i5
  channelName: My Channel
  customerName: Test Customer
  endpoint: https://replicated.app.modified
  entitlements:
    bool_field:
      title: Bool Field
      value: true
      valueType: Boolean
    expires_at:
      description: License Expiration
      title: Expiration
      value: "2030-07-27T00:00:00Z"
      valueType: String
    hidden_field:
      isHidden: true
      title: Hidden Field
      value: this is secret
      valueType: String
    int_field:
      title: Int Field
      value: 123
      valueType: Integer
    string_field:
      title: StringField
      value: single line text
      valueType: String
    text_field:
      title: Text Field
      value: |-
        multi
        line
        text
      valueType: Text
  isAirgapSupported: true
  isGitOpsSupported: true
  isSnapshotSupported: true
  licenseID: 1vusOokxAVp1tkRGuyxnF23PJcq
  licenseSequence: 7
  licenseType: prod
  signature: eyJsaWNlbnNlRGF0YSI6ImV5SmhjR2xXWlhKemFXOXVJam9pYTI5MGN5NXBieTkyTVdKbGRHRXhJaXdpYTJsdVpDSTZJa3hwWTJWdWMyVWlMQ0p0WlhSaFpHRjBZU0k2ZXlKdVlXMWxJam9pZEdWemRHTjFjM1J2YldWeUluMHNJbk53WldNaU9uc2liR2xqWlc1elpVbEVJam9pTVhaMWMwOXZhM2hCVm5BeGRHdFNSM1Y1ZUc1R01qTlFTbU54SWl3aWJHbGpaVzV6WlZSNWNHVWlPaUp3Y205a0lpd2lZM1Z6ZEc5dFpYSk9ZVzFsSWpvaVZHVnpkQ0JEZFhOMGIyMWxjaUlzSW1Gd2NGTnNkV2NpT2lKdGVTMWhjSEFpTENKamFHRnVibVZzU1VRaU9pSXhkblZ6U1ZsYVRFRldlRTFITm5FM05qQlBTbTFTUzJvMWFUVWlMQ0pqYUdGdWJtVnNUbUZ0WlNJNklrMTVJRU5vWVc1dVpXd2lMQ0pzYVdObGJuTmxVMlZ4ZFdWdVkyVWlPamNzSW1WdVpIQnZhVzUwSWpvaWFIUjBjSE02THk5eVpYQnNhV05oZEdWa0xtRndjQ0lzSW1WdWRHbDBiR1Z0Wlc1MGN5STZleUppYjI5c1gyWnBaV3hrSWpwN0luUnBkR3hsSWpvaVFtOXZiQ0JHYVdWc1pDSXNJblpoYkhWbElqcDBjblZsTENKMllXeDFaVlI1Y0dVaU9pSkNiMjlzWldGdUluMHNJbVY0Y0dseVpYTmZZWFFpT25zaWRHbDBiR1VpT2lKRmVIQnBjbUYwYVc5dUlpd2laR1Z6WTNKcGNIUnBiMjRpT2lKTWFXTmxibk5sSUVWNGNHbHlZWFJwYjI0aUxDSjJZV3gxWlNJNklqSXdNekF0TURjdE1qZFVNREE2TURBNk1EQmFJaXdpZG1Gc2RXVlVlWEJsSWpvaVUzUnlhVzVuSW4wc0ltaHBaR1JsYmw5bWFXVnNaQ0k2ZXlKMGFYUnNaU0k2SWtocFpHUmxiaUJHYVdWc1pDSXNJblpoYkhWbElqb2lkR2hwY3lCcGN5QnpaV055WlhRaUxDSjJZV3gxWlZSNWNHVWlPaUpUZEhKcGJtY2lMQ0pwYzBocFpHUmxiaUk2ZEhKMVpYMHNJbWx1ZEY5bWFXVnNaQ0k2ZXlKMGFYUnNaU0k2SWtsdWRDQkdhV1ZzWkNJc0luWmhiSFZsSWpveE1qTXNJblpoYkhWbFZIbHdaU0k2SWtsdWRHVm5aWElpZlN3aWMzUnlhVzVuWDJacFpXeGtJanA3SW5ScGRHeGxJam9pVTNSeWFXNW5SbWxsYkdRaUxDSjJZV3gxWlNJNkluTnBibWRzWlNCc2FXNWxJSFJsZUhRaUxDSjJZV3gxWlZSNWNHVWlPaUpUZEhKcGJtY2lmU3dpZEdWNGRGOW1hV1ZzWkNJNmV5SjBhWFJzWlNJNklsUmxlSFFnUm1sbGJHUWlMQ0oyWVd4MVpTSTZJbTExYkhScFhHNXNhVzVsWEc1MFpYaDBJaXdpZG1Gc2RXVlVlWEJsSWpvaVZHVjRkQ0o5ZlN3aWFYTkJhWEpuWVhCVGRYQndiM0owWldRaU9uUnlkV1VzSW1selIybDBUM0J6VTNWd2NHOXlkR1ZrSWpwMGNuVmxMQ0pwYzFOdVlYQnphRzkwVTNWd2NHOXlkR1ZrSWpwMGNuVmxmWDA9IiwiaW5uZXJTaWduYXR1cmUiOiJleUpzYVdObGJuTmxVMmxuYm1GMGRYSmxJam9pYUhneE1XTXZUR1ozUTNoVE5YRmtRWEJGU1hGdVRrMU9NMHBLYTJzNFZHZFhSVVpzVDFKVlJ6UjJjR1YzZEZoV1YzbG1lamRZY0hBd1ExazJZamRyUVRSS2N6TklhR3d3YkZJMFdUQTFMemN2UVVkQ2FEZFZNSGczUkhaTVozUXpVM00wYm5GTFZTdFhXRXBTVHpKWVFVRnZSME4xZFRWR1RGcHJRVWhYY1RSUVFtMXphSFY2Y1ZsdmNucHhlbGhGWVZWVlpFUlVkVXhDTW1nNWFIZ3dXRWhQUmxwUk16bHVkbTlPUjJaT2R5OTRTVmRaZEhSUGRYZHZhMncyTVZsb1JVeFZlRmQxU1ZSRmMwTlVhM2xtTVRNd09IazVSbFJzWlRKeVYyZEVlSEZNYTBSUFNXVXlPRWwzUzJSQkwySXdWVUl5VEZGbVRWcHdWemwyUTNCSkwybHlWek5uYmpaeU5WWjNWMjB2U1dweWJtNDNSelJrVmpadVYzcFRkMGhQUTJSdWEwMTRNRXQ1VVVOa0wxQjFaWEpUYjNSdVEwOXRTMDEzWlRSTGJqaERkMU5YVVRRNGRURkRNbTFpV1VzeGRYTlpOM1YzUFQwaUxDSndkV0pzYVdOTFpYa2lPaUl0TFMwdExVSkZSMGxPSUZCVlFreEpReUJMUlZrdExTMHRMVnh1VFVsSlFrbHFRVTVDWjJ0eGFHdHBSemwzTUVKQlVVVkdRVUZQUTBGUk9FRk5TVWxDUTJkTFEwRlJSVUZ6TkhKdlVIcDFhV1JNZVhOMmIxWTJkemxhTkZ4dVdHRmliME5tWTJNeGFHZFZhQ3N3V1VkS2NFNURSVXhyTjBaTFF5OTJhemR6ZERsR05tY3dUMjlrU0VSbGVYZFJXa2hLZFU1TVpsUnNRbEJHUTJOaU5seHVObTlzVEZOeWNGQTRjbFUzU0d4SGJsRkVSMFJNYVhkS1EyaGtSRGRVVUdSM2FXdHBkMHRGY201aldqaEdaalZsU25vd2RETmlUWFpyVDJaVVluSkJiRnh1WWtGQ1kwbzVNVmxVT1hKdVVXOXFkVWN4UldKUVRqaEZWblI2TWxZNE5IZHViR2Q0TUhCd2JEVjRPSFpOYlhwcE1ISnVibEZVV1VGamJ6WnFhMnBJTTF4dVRuTlVkWE4xUzFkdlJGUjVNWE5yZGtSUk9IbEJZV0ptWTNNME4zWnNRazAwU0RGT1JFNHZSSFJhWWxZdllubDJia0o2YkM4eFZrVnpURmRqWlZWcFRGeHVSWEYxT0VkeWF5dFFVRGQyUkdSd2JFUjNjWFpQV2t4RmRYazNkamhuUm01U09WUlVSV3ByTlVvNWRuWlVTR2RtU25VemVubEVPR2xLWTBSRE5YcHFPVnh1YjFGSlJFRlJRVUpjYmkwdExTMHRSVTVFSUZCVlFreEpReUJMUlZrdExTMHRMVnh1SWl3aWEyVjVVMmxuYm1GMGRYSmxJam9pWlhsS2VtRlhaSFZaV0ZJeFkyMVZhVTlwU2pCUldIQjJXVE5LVms1NmFGaFNSMlJzVVRKb2NtTklXa1ZVVlRsRldqQktXVTFGUmtaVFJFNUZVMGhLYkUxclRUTkxNSEJFVkROR2VGTnROVVJVVlRWVlltMDFiVnBGUm5sWldIQjZaRVJqTVZaSGFFeFBXRUpVVWtacmRrd3diek5aTUZaSlVteFdWRXd5T1VoV1JXeHNWa1ZPTUZSSE1WWlJNR04zVkd4R2JGa3pTblJUUm1zMFZVWk9hMVpWU2pCVU1WbDNZbXQwY0ZSclZuQmpia0poVFZjNWFtSldiSEZaYTNob1UyeHNWV0pGUmtWWGJVWnZWakZLVUZkcWJGSmhXRVp1V2xkb1EyRnVRak5TUjNNd1lWWkpOVTVXVmxkV1ZUVnlUMGhLYjFsVlRYbGhiVGcwVjBkYWVGbHFWbFppYlhoeFpFWkZkMDU1Y3pCaFZsSkpWRVpPTm1WRk1IcGxWWFJ2VFVaR1ZtRXdWVFJSVnpsSFVsaEtVRTFZUmxCU01WcFJVMVJDTmxsV2FIcFdWWEJ0WTBSU2JFMVVRazlPVjNSU1ZucFdUMU5XWTNaU1ZYUkZVMGhzYlU5VmJGaGtNMUl3WTFWc1lXTlhSakJTYTA1RVlVWmtjbUo2VmtSU00wSllUREkxUmsxWVl6SmxWM1JKVlZoQk1sVXhTbEppU0Zwd1VrVXdNRlpFVWt0VU1rWnNVVmQwYzFSV1VrMVVWV055V1RCYVRHSXpaRTlUVm05NVlraE9SR1JzVG5aUmFrWmFaVmRPVGxOVlNteGFiRXB1Wld0U2RVMHhSVGxRVTBselNXMWtjMkl5U21oaVJYUnNaVlZzYTBscWIybFpiVkpzV2xSVk1rNVVXWGRaTWxwcFRrUk9hazlYU1hsUFIwcHRUMVJvYkZsWFRtaGFiVVV5VGtSWmFXWlJQVDBpZlE9PSJ9
`,
			wantErr:    true,
			wantErrMsg: `"endpoint" field has changed`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := require.New(t)

			license, err := LoadLicenseFromBytes([]byte(tt.license))
			req.NoError(err)

			_, err = VerifySignature(license)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifySignature() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && tt.wantErrMsg != err.Error() {
				t.Errorf("VerifySignature() error message = %v, wantErrMsg %v", err.Error(), tt.wantErrMsg)
				return
			}
		})
	}
}
