# frozen_string_literal: true

RSpec.describe ParseMonthService do
  subject(:service) { described_class.new(params) }

  let(:params) { { month: '2021-03' } }

  describe '#month' do
    its(:month) { is_expected.to eq Month.new(2021, 3) }

    context 'with default date' do
      let(:params) { {} }

      before { travel_to Date.new(2021, 4, 1) }

      its(:month) { is_expected.to eq Month.new(2021, 4) }
    end
  end

  describe '#call' do
    let(:month) { Month.now }

    before { allow(service).to receive(:month).and_return(month) }

    its(:call) { is_expected.to eq month }
  end

  describe '.call' do
    before do
      #
      # described_class.new(params).call
      #
      allow(described_class).to receive(:new).with(params) do
        double.tap do |a|
          allow(a).to receive(:call)
        end
      end
    end

    it { expect { described_class.call(params) }.not_to raise_error }
  end
end
