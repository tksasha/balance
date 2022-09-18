# frozen_string_literal: true

RSpec.describe Currency do
  describe '.parse' do
    subject { described_class.parse(currency) }

    context 'with "usd"' do
      let(:currency) { 'usd' }

      it { is_expected.to eq 'usd' }
    end

    context 'with "USD"' do
      let(:currency) { 'USD' }

      it { is_expected.to eq 'usd' }
    end

    context 'with `nil`' do
      let(:currency) { nil }

      it { is_expected.to eq 'uah' }
    end

    context 'with "unsupported currency"' do
      let(:currency) { 'unsupported currency' }

      it { is_expected.to eq 'uah' }
    end

    context 'with "uah"' do
      let(:currency) { 'uah' }

      it { is_expected.to eq 'uah' }
    end
  end

  describe 'Currency' do
    subject { Currency('USD') }

    it { is_expected.to eq 'usd' }
  end
end
