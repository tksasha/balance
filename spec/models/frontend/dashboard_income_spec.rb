# frozen_string_literal: true

RSpec.describe Frontend::Dashboard do
  subject { described_class.new(currency:, year: 2022, month: 10) }

  describe '#income' do
    let(:income) { create(:category, income: true) }

    let(:expense) { create(:category, income: false) }

    before do
      create(:item, category: income, currency: 'uah', sum: 1.1)
      create(:item, category: income, currency: 'usd', sum: 2.2)
      create(:item, category: income, currency: 'eur', sum: 3.3)

      create(:item, category: expense, currency: 'uah')
    end

    context 'when currency is `uah`' do
      let(:currency) { 'uah' }

      before { create(:item, category: income, currency:, sum: 5.5) }

      it { expect(subject.income).to eq 6.6 }
    end

    context 'when currency is `usd`' do
      let(:currency) { 'usd' }

      before { create(:item, category: income, currency:, sum: 5.5) }

      it { expect(subject.income).to eq 7.7 }
    end

    context 'when currency is `eur`' do
      let(:currency) { 'eur' }

      before { create(:item, category: income, currency:, sum: 5.5) }

      it { expect(subject.income).to eq 8.8 }
    end
  end
end
