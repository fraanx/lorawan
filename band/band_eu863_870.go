package band

import (
	"time"

	"github.com/fraanx/lorawan"
)

type eu863Band struct {
	band
}

func (b *eu863Band) Name() string {
	return "EU868"
}

func (b *eu863Band) GetDefaults() Defaults {
	return Defaults{
		RX2Frequency:     869525000,
		RX2DataRate:      0,
		ReceiveDelay1:    time.Second,
		ReceiveDelay2:    time.Second * 2,
		JoinAcceptDelay1: time.Second * 5,
		JoinAcceptDelay2: time.Second * 6,
	}
}

func (b *eu863Band) GetDownlinkTXPower(freq uint32) int {
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

func (b *eu863Band) GetDefaultMaxUplinkEIRP() float32 {
	return 16
}

func (b *eu863Band) GetPingSlotFrequency(lorawan.DevAddr, time.Duration) (uint32, error) {
	return 869525000, nil
}

func (b *eu863Band) GetRX1ChannelIndexForUplinkChannelIndex(uplinkChannel int) (int, error) {
	return uplinkChannel, nil
}

func (b *eu863Band) GetRX1FrequencyForUplinkFrequency(uplinkFrequency uint32) (uint32, error) {
	return uplinkFrequency, nil
}

func (b *eu863Band) ImplementsTXParamSetup(protocolVersion string) bool {
	return false
}

func newEU863Band(repeatedCompatible bool) (Band, error) {
	b := eu863Band{
		band: band{
			supportsExtraChannels: true,
			cFListMinDR:           0,
			cFListMaxDR:           5,
			dataRates: map[int]DataRate{
				0:  {Modulation: LoRaModulation, SpreadFactor: 12, Bandwidth: 125, uplink: true, downlink: true},
				1:  {Modulation: LoRaModulation, SpreadFactor: 11, Bandwidth: 125, uplink: true, downlink: true},
				2:  {Modulation: LoRaModulation, SpreadFactor: 10, Bandwidth: 125, uplink: true, downlink: true},
				3:  {Modulation: LoRaModulation, SpreadFactor: 9, Bandwidth: 125, uplink: true, downlink: true},
				4:  {Modulation: LoRaModulation, SpreadFactor: 8, Bandwidth: 125, uplink: true, downlink: true},
				5:  {Modulation: LoRaModulation, SpreadFactor: 7, Bandwidth: 125, uplink: true, downlink: true},
				6:  {Modulation: LoRaModulation, SpreadFactor: 7, Bandwidth: 250, uplink: true, downlink: true},
				7:  {Modulation: FSKModulation, BitRate: 50000, uplink: true, downlink: true},
				8:  {Modulation: LRFHSSModulation, CodingRate: "1/3", OccupiedChannelWidth: 137000, uplink: true, downlink: false},
				9:  {Modulation: LRFHSSModulation, CodingRate: "4/6", OccupiedChannelWidth: 137000, uplink: true, downlink: false},
				10: {Modulation: LRFHSSModulation, CodingRate: "1/3", OccupiedChannelWidth: 336000, uplink: true, downlink: false},
				11: {Modulation: LRFHSSModulation, CodingRate: "4/6", OccupiedChannelWidth: 336000, uplink: true, downlink: false},
			},
			rx1DataRateTable: map[int][]int{
				0:  {0, 0, 0, 0, 0, 0},
				1:  {1, 0, 0, 0, 0, 0},
				2:  {2, 1, 0, 0, 0, 0},
				3:  {3, 2, 1, 0, 0, 0},
				4:  {4, 3, 2, 1, 0, 0},
				5:  {5, 4, 3, 2, 1, 0},
				6:  {6, 5, 4, 3, 2, 1},
				7:  {7, 6, 5, 4, 3, 2},
				8:  {1, 0, 0, 0, 0, 0},
				9:  {2, 1, 0, 0, 0, 0},
				10: {1, 0, 0, 0, 0, 0},
				11: {2, 1, 0, 0, 0, 0},
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
				{Frequency: 868100000, MinDR: 0, MaxDR: 5, enabled: true},
				{Frequency: 868300000, MinDR: 0, MaxDR: 5, enabled: true},
				{Frequency: 868500000, MinDR: 0, MaxDR: 5, enabled: true},
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
					6: {M: 230, N: 222},
					7: {M: 230, N: 222},
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
					6: {M: 230, N: 222},
					7: {M: 230, N: 222},
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
					6: {M: 230, N: 222},
					7: {M: 230, N: 222},
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
					6: {M: 230, N: 222},
					7: {M: 230, N: 222},
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
					6: {M: 230, N: 222},
					7: {M: 230, N: 222},
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
					6: {M: 230, N: 222},
					7: {M: 230, N: 222},
				},
				RegParamRevRP002_1_0_1: map[int]MaxPayloadSize{
					0: {M: 59, N: 51},
					1: {M: 59, N: 51},
					2: {M: 59, N: 51},
					3: {M: 123, N: 115},
					4: {M: 230, N: 222},
					5: {M: 230, N: 222},
					6: {M: 230, N: 222},
					7: {M: 230, N: 222},
				},
				latest: map[int]MaxPayloadSize{ // RP002-1.0.2, RP002-1.0.3
					0:  {M: 59, N: 51},
					1:  {M: 59, N: 51},
					2:  {M: 59, N: 51},
					3:  {M: 123, N: 115},
					4:  {M: 230, N: 222},
					5:  {M: 230, N: 222},
					6:  {M: 230, N: 222},
					7:  {M: 230, N: 222},
					8:  {M: 58, N: 50},
					9:  {M: 123, N: 115},
					10: {M: 58, N: 50},
					11: {M: 123, N: 115},
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
					6: {M: 250, N: 242},
					7: {M: 250, N: 242},
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
					6: {M: 250, N: 242},
					7: {M: 250, N: 242},
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
					6: {M: 250, N: 242},
					7: {M: 250, N: 242},
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
					6: {M: 250, N: 242},
					7: {M: 250, N: 242},
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
					6: {M: 250, N: 242},
					7: {M: 250, N: 242},
				},
				RegParamRevRP002_1_0_1: map[int]MaxPayloadSize{
					0: {M: 59, N: 51},
					1: {M: 59, N: 51},
					2: {M: 59, N: 51},
					3: {M: 123, N: 115},
					4: {M: 250, N: 242},
					5: {M: 250, N: 242},
					6: {M: 250, N: 242},
					7: {M: 250, N: 242},
				},
				latest: map[int]MaxPayloadSize{ // RP002-1.0.2, RP002-1.0.3
					0:  {M: 59, N: 51},
					1:  {M: 59, N: 51},
					2:  {M: 59, N: 51},
					3:  {M: 123, N: 115},
					4:  {M: 250, N: 242},
					5:  {M: 250, N: 242},
					6:  {M: 250, N: 242},
					7:  {M: 250, N: 242},
					8:  {M: 58, N: 50},
					9:  {M: 123, N: 115},
					10: {M: 58, N: 50},
					11: {M: 123, N: 115},
				},
			},
		}
	}

	return &b, nil
}
