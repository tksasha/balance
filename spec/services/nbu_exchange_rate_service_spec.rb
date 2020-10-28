# frozen_string_literal: true

RSpec.describe NbuExchangeRateService do
  let(:date) { Date.new 2019, 12, 31 }

  let(:url) { 'https://www.bank.gov.ua/markets/exchangerates?date=31.12.2019&period=daily' }

  subject { described_class.new date }

  describe '#date' do
    its(:date) { should eq '31.12.2019' }

    context do
      before { travel_to Date.new 2019, 12, 14 }

      after { travel_back }

      subject { described_class.new }

      its(:date) { should eq '14.12.2019' }
    end
  end

  describe '#rates' do
    let(:data) { File.open Rails.root.join('spec/support/nbu.html') }

    before { expect(URI).to receive_message_chain(:parse, :open).with(url).with(no_args).and_return(data) }

    its(:rates) { should eq usd: 24.538038, rub: 0.38207 }
  end

  describe '.rates' do
    let(:date) { Date.today }

    before do
      #
      # described_class.new(date).rates -> :rates
      #
      allow(described_class).to receive(:new).with(date) do
        double.tap do |a|
          allow(a).to receive(:rates).and_return(:rates)
        end
      end
    end

    subject { described_class.rates date }

    it { should eq :rates }
  end
end
