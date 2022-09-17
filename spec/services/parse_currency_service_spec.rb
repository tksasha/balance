# frozen_string_literal: true

RSpec.describe ParseCurrencyService do
  subject { described_class.new currency }

  describe '#currency' do
    context do
      let(:currency) { 'usd' }

      its(:currency) { is_expected.to eq 'usd' }
    end

    context do
      let(:currency) { 'USD' }

      its(:currency) { is_expected.to eq 'usd' }
    end

    context do
      let(:currency) { nil }

      its(:currency) { is_expected.to eq 'uah' }
    end

    context do
      let(:currency) { 'unsupported currency' }

      its(:currency) { is_expected.to eq 'uah' }
    end

    context do
      let(:currency) { 'uah' }

      its(:currency) { is_expected.to eq 'uah' }
    end

    context do
      let(:currency) { 'rub' }

      its(:currency) { is_expected.to eq 'rub' }
    end
  end

  describe '#call' do
    let(:currency) { double }

    before { allow(subject).to receive(:currency).and_return(currency) }

    its(:call) { is_expected.to eq currency }
  end

  describe '.call' do
    subject { described_class.call('usd') }

    before do
      #
      # described_class.new('usd').call
      #
      expect(described_class).to receive(:new).with('usd') do
        double.tap do |a|
          expect(a).to receive(:call)
        end
      end
    end

    it { expect { subject }.not_to raise_error }
  end
end
