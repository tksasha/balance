# frozen_string_literal: true

RSpec.describe Frontend::Dashboard::Cashes do
  describe '#sum' do
    subject { described_class.new(currency:).sum }

    before do
      create(:cash, :uah, sum: 1.01)
      create(:cash, :uah, sum: 1.02)

      create(:cash, :usd, sum: 2.03)
      create(:cash, :usd, sum: 2.04)

      create(:cash, :eur, sum: 3.05)
      create(:cash, :eur, sum: 3.06)
    end

    context 'when `uah`' do
      let(:currency) { 'uah' }

      it { is_expected.to be_a(BigDecimal) }

      it { is_expected.to eq 2.03 }
    end

    context 'when `usd`' do
      let(:currency) { 'usd' }

      it { is_expected.to be_a(BigDecimal) }

      it { is_expected.to eq 4.07 }
    end

    context 'when `eur`' do
      let(:currency) { 'eur' }

      it { is_expected.to be_a(BigDecimal) }

      it { is_expected.to eq 6.11 }
    end
  end
end
