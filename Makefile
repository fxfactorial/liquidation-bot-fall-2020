liq_bot := liq_bot
arb_gainer := arb_gainer
pair_alerter := pair_alerter
emerg_swap := must_swap
scam_pairs := rugpull_pairs
evm_runner := evm_runner
origin := origin

.PHONY: liq_bot

all: new_pair_exe arb_exe emerg_swap pair_scammer custom_evm_runner liq_bot origin

ousd:
	go build -o ${origin} cmd/origin/*.go

liq_bot:
	go build -o ${liq_bot} cmd/liq_bot/*.go

custom_evm_runner:
	go build -o ${evm_runner} cmd/eval_evm/*.go

pair_scammer:
	go build -o ${scam_pairs} cmd/rugpull/*.go

pair_scammer_ropsten:
	go run cmd/rugpull/*.go -client_dial 'ws://127.0.0.1:3546' -manager_addr $(MANAGER_ADDR) -ropsten

pair_scammer_run:
	go run cmd/rugpull/*.go -client_dial 'ws://127.0.0.1:8545' -manager_addr $(MANAGER_ADDR) -ropsten

new_pair_exe:
	go build -o ${pair_alerter} cmd/pair_alert/*.go

arb_exe:
	go build -o ${arb_gainer} cmd/arb_gain/*.go

emerg_swap:
	go build -o ${emerg_swap} cmd/emergency_swap/*.go
