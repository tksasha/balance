# frozen_string_literal: true

RSpec.describe Frontend::Reports::Cashes do
  describe '.call' do
    subject { described_class.call(currency:) }

    before do
      create(:cash, currency: 'uah', supercategory: 'cash', id: 1, name: 'Wallet', sum: 11.11)
      create(:cash, currency: 'usd', supercategory: 'cash', id: 2, name: 'Wallet', sum: 22.22)
      create(:cash, currency: 'eur', supercategory: 'cash', id: 3, name: 'Wallet', sum: 33.33)
      create(:cash, currency: 'uah', supercategory: 'deposits', id: 4, name: 'Monobank', sum: 44.44)
      create(:cash, currency: 'usd', supercategory: 'deposits', id: 5, name: 'Monobank', sum: 55.55)
      create(:cash, currency: 'eur', supercategory: 'deposits', id: 6, name: 'Monobank', sum: 66.66)
      create(:cash, currency: 'uah', supercategory: 'bonds', id: 7, name: 'ICU Trade', sum: 77.77)
      create(:cash, currency: 'usd', supercategory: 'bonds', id: 8, name: 'ICU Trade', sum: 88.88)
      create(:cash, currency: 'eur', supercategory: 'bonds', id: 9, name: 'ICU Trade', sum: 99.99)
    end

    context 'when currency is `uah`' do
      let(:currency) { 'uah' }

      let(:cashes) do
        {
          'cash' => [['cash', 1, 'Wallet', 11.11]],
          'deposits' => [['deposits', 4, 'Monobank', 44.44]],
          'bonds' => [['bonds', 7, 'ICU Trade', 77.77]]
        }
      end

      it { is_expected.to eq cashes }
    end

    context 'when currency is `usd`' do
      let(:currency) { 'usd' }

      let(:cashes) do
        {
          'cash' => [['cash', 2, 'Wallet', 22.22]],
          'deposits' => [['deposits', 5, 'Monobank', 55.55]],
          'bonds' => [['bonds', 8, 'ICU Trade', 88.88]]
        }
      end

      it { is_expected.to eq cashes }
    end

    context 'when currency is `eur`' do
      let(:currency) { 'eur' }

      let(:cashes) do
        {
          'cash' => [['cash', 3, 'Wallet', 33.33]],
          'deposits' => [['deposits', 6, 'Monobank', 66.66]],
          'bonds' => [['bonds', 9, 'ICU Trade', 99.99]]
        }
      end

      it { is_expected.to eq cashes }
    end
  end
end
