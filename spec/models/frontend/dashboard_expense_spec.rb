# frozen_string_literal: true

RSpec.describe Frontend::Dashboard do
  subject { described_class.new(currency:, year: 2022, month: 11) }

  describe '#expense' do
    before do
      create(:item, :expense, :uah, sum: 1.1)
      create(:item, :expense, :usd, sum: 2.2)
      create(:item, :expense, :eur, sum: 3.3)

      create(:item, :income, :uah, sum: 4.4)
    end

    context 'when currency is `uah`' do
      let(:currency) { 'uah' }

      before { create(:item, :expense, currency:, sum: 5.5) }

      it { expect(subject.expense).to eq 6.6 }
    end

    context 'when currency is `usd`' do
      let(:currency) { 'usd' }

      before { create(:item, :expense, currency:, sum: 5.5) }

      it { expect(subject.expense).to eq 7.7 }
    end

    context 'when currency is `eur`' do
      let(:currency) { 'eur' }

      before { create(:item, :expense, currency:, sum: 5.5) }

      it { expect(subject.expense).to eq 8.8 }
    end
  end
end
