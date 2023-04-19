package band

import (
	"time"

	"github.com/fraanx/lorawan"
)

type ss001Band struct {
	band
}

func (b *ss001Band) Name() string {
	return "SS001"
}

func (b *ss001Band) GetDefaults() Defaults {
	return Defaults{
		RX2Frequency:     869525000,
		RX2DataRate:      0,
		ReceiveDelay1:    time.Second,
		ReceiveDelay2:    time.Second * 2,
		JoinAcceptDelay1: time.Second * 5,
		JoinAcceptDelay2: time.Second * 6,
	}
}

func (b *ss001Band) GetDownlinkTXPower(freq uint32) int {
	// NOTE: as there are currently no further boundary checks on the frequency, this check is sufficient.
	// TODO: However, there should be some mechanism, that checks the frequency for compliance to regulations.
	if 863000000 <= freq && freq < 869200000 {
		return 14 //25mW
	} else if 869400000 <= freq && freq < 869650000 {
		return 27 //500mW
	} else {
		return 14 // Default case
	}
}

func (b *ss001Band) GetDefaultMaxUplinkEIRP() float32 {
	return 16
}

func (b *ss001Band) GetPingSlotFrequency(lorawan.DevAddr, time.Duration) (uint32, error) {
	return 869525000, nil
}

func (b *ss001Band) GetRX1ChannelIndexForUplinkChannelIndex(uplinkChannel int) (int, error) {
	return uplinkChannel, nil
}

func (b *ss001Band) GetRX1FrequencyForUplinkFrequency(uplinkFrequency uint32) (uint32, error) {
	return uplinkFrequency, nil
}

func (b *ss001Band) ImplementsTXParamSetup(protocolVersion string) bool {
	return false
}

func newSS001Band(repeatedCompatible bool) (Band, error) {
	b := ss001Band{
		band: band{
			supportsExtraChannels: true,
			cFListMinDR:           0,
			cFListMaxDR:           5,
			dataRates: map[int]DataRate{
				0:  {Modulation: LoRaModulation, SpreadFactor: 64, Bandwidth: 240, uplink: true, downlink: false},
				1:  {Modulation: LoRaModulation, SpreadFactor: 128, Bandwidth: 240, uplink: true, downlink: false},
				2:  {Modulation: LoRaModulation, SpreadFactor: 256, Bandwidth: 240, uplink: true, downlink: false},
				3:  {Modulation: LoRaModulation, SpreadFactor: 64, Bandwidth: 240, uplink: true, downlink: false},
				4:  {Modulation: LoRaModulation, SpreadFactor: 128, Bandwidth: 240, uplink: true, downlink: false},
				5:  {Modulation: LoRaModulation, SpreadFactor: 256, Bandwidth: 240, uplink: true, downlink: false},
				11:  {Modulation: LoRaModulation, SpreadFactor: 16, Bandwidth: 2400, uplink: false, downlink: true},
			},
			rx1DataRateTable: map[int][]int{
				0:  {11},
				1:  {11},
				2:  {11},
				3:  {11},
				4:  {11},
				5:  {11},
			},

			txPowerOffsets: []int{
				0,
				-2,
				-4,
				-6,
				-8,
				-10,
				-12,
				-14,
			},
			uplinkChannels: []Channel{
				{Frequency: 868100000, MinDR: 0, MaxDR: 5, enabled: true},
				{Frequency: 868300000, MinDR: 0, MaxDR: 5, enabled: true},
				{Frequency: 868500000, MinDR: 0, MaxDR: 5, enabled: true},
			},
			downlinkChannels: []Channel{
				{Frequency: 868100000, MinDR: 11, MaxDR: 11, enabled: true},
				{Frequency: 868300000, MinDR: 11, MaxDR: 11, enabled: true},
				{Frequency: 868500000, MinDR: 11, MaxDR: 11, enabled: true},
			},
		},
	}

	if repeatedCompatible {
		b.band.maxPayloadSizePerDR = map[string]map[string]map[int]MaxPayloadSize{
			LoRaWAN_1_0_0: map[string]map[int]MaxPayloadSize{
				latest: map[int]MaxPayloadSize{ // LoRaWAN 1.0.0
					0: {M: 59, N: 51},
					1: {M: 59, N: 51},
					2: {M: 59, N: 51},
					3: {M: 123, N: 115},
					4: {M: 230, N: 222},
					5: {M: 230, N: 222},
					11: {M: 230, N: 222},
				},
			},
			LoRaWAN_1_0_1: map[string]map[int]MaxPayloadSize{
				latest: map[int]MaxPayloadSize{ // LoRaWAN 1.0.1
					0: {M: 59, N: 51},
					1: {M: 59, N: 51},
					2: {M: 59, N: 51},
					3: {M: 123, N: 115},
					4: {M: 230, N: 222},
					5: {M: 230, N: 222},
					11: {M: 230, N: 222},
				},
			},
			LoRaWAN_1_0_2: map[string]map[int]MaxPayloadSize{
				latest: map[int]MaxPayloadSize{ // LoRaWAN 1.0.2A, 1.0.2B
					0: {M: 59, N: 51},
					1: {M: 59, N: 51},
					2: {M: 59, N: 51},
					3: {M: 123, N: 115},
					4: {M: 230, N: 222},
					5: {M: 230, N: 222},
					11: {M: 230, N: 222},
				},
			},
			LoRaWAN_1_0_3: map[string]map[int]MaxPayloadSize{
				latest: map[int]MaxPayloadSize{ // LoRaWAN 1.0.3A
					0: {M: 59, N: 51},
					1: {M: 59, N: 51},
					2: {M: 59, N: 51},
					3: {M: 123, N: 115},
					4: {M: 230, N: 222},
					5: {M: 230, N: 222},
					11: {M: 230, N: 222},
				},
			},
			LoRaWAN_1_1_0: map[string]map[int]MaxPayloadSize{
				latest: map[int]MaxPayloadSize{ // LoRaWAN 1.1.0A, 1.1.0B
					0: {M: 59, N: 51},
					1: {M: 59, N: 51},
					2: {M: 59, N: 51},
					3: {M: 123, N: 115},
					4: {M: 230, N: 222},
					5: {M: 230, N: 222},
					11: {M: 230, N: 222},
				},
			},
			latest: map[string]map[int]MaxPayloadSize{
				RegParamRevRP002_1_0_0: map[int]MaxPayloadSize{
					0: {M: 59, N: 51},
					1: {M: 59, N: 51},
					2: {M: 59, N: 51},
					3: {M: 123, N: 115},
					4: {M: 230, N: 222},
					5: {M: 230, N: 222},
					11: {M: 230, N: 222},
				},
				RegParamRevRP002_1_0_1: map[int]MaxPayloadSize{
					0: {M: 59, N: 51},
					1: {M: 59, N: 51},
					2: {M: 59, N: 51},
					3: {M: 123, N: 115},
					4: {M: 230, N: 222},
					5: {M: 230, N: 222},
					11: {M: 230, N: 222},
				},
				latest: map[int]MaxPayloadSize{ // RP002-1.0.2, RP002-1.0.3
					0:  {M: 59, N: 51},
					1:  {M: 59, N: 51},
					2:  {M: 59, N: 51},
					3:  {M: 123, N: 115},
					4:  {M: 230, N: 222},
					5:  {M: 230, N: 222},
					11: {M: 230, N: 222},
				},
			},
		}
	} else {
		b.band.maxPayloadSizePerDR = map[string]map[string]map[int]MaxPayloadSize{
			LoRaWAN_1_0_0: map[string]map[int]MaxPayloadSize{
				latest: map[int]MaxPayloadSize{ // LoRaWAN 1.0.0
					0: {M: 59, N: 51},
					1: {M: 59, N: 51},
					2: {M: 59, N: 51},
					3: {M: 123, N: 115},
					4: {M: 250, N: 242},
					5: {M: 250, N: 242},
					11: {M: 230, N: 222},
				},
			},
			LoRaWAN_1_0_1: map[string]map[int]MaxPayloadSize{
				latest: map[int]MaxPayloadSize{ // LoRaWAN 1.0.1
					0: {M: 59, N: 51},
					1: {M: 59, N: 51},
					2: {M: 59, N: 51},
					3: {M: 123, N: 115},
					4: {M: 250, N: 242},
					5: {M: 250, N: 242},
					11: {M: 230, N: 222},
				},
			},
			LoRaWAN_1_0_2: map[string]map[int]MaxPayloadSize{
				latest: map[int]MaxPayloadSize{ // LoRaWAN 1.0.2A, 1.0.2B
					0: {M: 59, N: 51},
					1: {M: 59, N: 51},
					2: {M: 59, N: 51},
					3: {M: 123, N: 115},
					4: {M: 250, N: 242},
					5: {M: 250, N: 242},
					11: {M: 230, N: 222},
				},
			},
			LoRaWAN_1_0_3: map[string]map[int]MaxPayloadSize{
				latest: map[int]MaxPayloadSize{ // LoRaWAN 1.0.3A
					0: {M: 59, N: 51},
					1: {M: 59, N: 51},
					2: {M: 59, N: 51},
					3: {M: 123, N: 115},
					4: {M: 250, N: 242},
					5: {M: 250, N: 242},
					11: {M: 230, N: 222},
				},
			},
			latest: map[string]map[int]MaxPayloadSize{
				RegParamRevRP002_1_0_0: map[int]MaxPayloadSize{
					0: {M: 59, N: 51},
					1: {M: 59, N: 51},
					2: {M: 59, N: 51},
					3: {M: 123, N: 115},
					4: {M: 250, N: 242},
					5: {M: 250, N: 242},
					11: {M: 230, N: 222},
				},
				RegParamRevRP002_1_0_1: map[int]MaxPayloadSize{
					0: {M: 59, N: 51},
					1: {M: 59, N: 51},
					2: {M: 59, N: 51},
					3: {M: 123, N: 115},
					4: {M: 250, N: 242},
					5: {M: 250, N: 242},
					11: {M: 230, N: 222},
				},
				latest: map[int]MaxPayloadSize{ // RP002-1.0.2, RP002-1.0.3
					0:  {M: 59, N: 51},
					1:  {M: 59, N: 51},
					2:  {M: 59, N: 51},
					3:  {M: 123, N: 115},
					4:  {M: 250, N: 242},
					5:  {M: 250, N: 242},
					11: {M: 230, N: 222},
				},
			},
		}
	}

	return &b, nil
}
