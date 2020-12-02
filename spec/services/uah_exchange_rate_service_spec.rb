# frozen_string_literal: true

RSpec.describe UahExchangeRateService do
  let(:date) { Date.new(2019, 12, 31) }

  subject { described_class.new date }

  describe '#date' do
    its(:date) { should eq date }

    context do
      let(:date) { Date.new(2019, 12, 1) }

      before { travel_to date }

      after { travel_back }

      subject { described_class.new }

      its(:date) { should eq date }
    end
  end

  describe '#rates' do
    context do
      before { subject.instance_variable_set :@rates, :rates }

      its(:rates) { should eq :rates }
    end

    context do
      before { allow(NbuExchangeRateService).to receive(:rates).with(date).and_return(:rates) }

      its(:rates) { should eq :rates }
    end
  end

  describe '#create_usd' do
    let(:rates) { { usd: 23.47, rub: 0.389 } }

    before { allow(subject).to receive(:rates).and_return(rates) }

    before { expect(ExchangeRate).to receive(:create).with(from: :uah, to: :usd, date: date, rate: 23.47) }

    it { expect { subject.send(:create_usd) }.not_to raise_error }
  end

  describe '#create_rub' do
    let(:rates) { { usd: 23.47, rub: 0.389 } }

    before { allow(subject).to receive(:rates).and_return(rates) }

    before { expect(ExchangeRate).to receive(:create).with(from: :uah, to: :rub, date: date, rate: 0.389) }

    it { expect { subject.send(:create_rub) }.not_to raise_error }
  end

  describe '#save' do
    before { expect(subject).to receive(:create_usd) }

    before { expect(subject).to receive(:create_rub) }

    it { expect { subject.save }.not_to raise_error }
  end

  describe '.create' do
    before do
      #
      # described_class.new(date).save
      #
      expect(described_class).to receive(:new).with(date) do
        double.tap do |a|
          expect(a).to receive(:save)
        end
      end
    end

    subject { described_class.create date }

    it { expect { subject }.not_to raise_error }
  end
end
