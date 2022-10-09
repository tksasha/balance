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

  # rubocop:disable RSpec/MultipleMemoizedHelpers
  describe '#all' do
    subject { described_class.new(currency:).all }

    let!(:cash_n1) { create(:cash, :cash, :uah) }
    let!(:cash_n2) { create(:cash, :cash, :usd) }
    let!(:cash_n3) { create(:cash, :cash, :eur) }

    let!(:cash_n4) { create(:cash, :bonds, :uah) }
    let!(:cash_n5) { create(:cash, :bonds, :usd) }
    let!(:cash_n6) { create(:cash, :bonds, :eur) }

    let!(:cash_n7) { create(:cash, :deposits, :uah) }
    let!(:cash_n8) { create(:cash, :deposits, :usd) }
    let!(:cash_n9) { create(:cash, :deposits, :eur) }

    context 'when currency is `uah`' do
      let(:currency) { 'uah' }

      let(:cashes) do
        {
          'cash' => [cash_n1],
          'bonds' => [cash_n4],
          'deposits' => [cash_n7]
        }
      end

      it { is_expected.to eq cashes }
    end

    context 'when currency is `usd`' do
      let(:currency) { 'usd' }

      let(:cashes) do
        {
          'cash' => [cash_n2],
          'bonds' => [cash_n5],
          'deposits' => [cash_n8]
        }
      end

      it { is_expected.to eq cashes }
    end

    context 'when currency is `eur`' do
      let(:currency) { 'eur' }

      let(:cashes) do
        {
          'cash' => [cash_n3],
          'bonds' => [cash_n6],
          'deposits' => [cash_n9]
        }
      end

      it { is_expected.to eq cashes }
    end
  end
  # rubocop:enable RSpec/MultipleMemoizedHelpers
end
