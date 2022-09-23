# frozen_string_literal: true

RSpec.describe Frontend::Dashboard do
  subject { described_class.new(params) }

  let(:params) { {} }

  describe '#expense' do
    before do
      create(:item, :expense, :uah, sum: 1.1)
      create(:item, :expense, :usd, sum: 2.2)
      create(:item, :expense, :eur, sum: 3.3)

      create(:item, :income, :uah, sum: 4.4)
    end

    context 'when currency is `uah`' do
      let(:params) { { currency: 'uah' } }

      before { create(:item, :expense, :uah, sum: 5.5) }

      it { expect(subject.expense).to eq 6.6 }
    end

    context 'when currency is `usd`' do
      let(:params) { { currency: 'usd' } }

      before { create(:item, :expense, :usd, sum: 5.5) }

      it { expect(subject.expense).to eq 7.7 }
    end

    context 'when currency is `eur`' do
      let(:params) { { currency: 'eur' } }

      before { create(:item, :expense, :eur, sum: 5.5) }

      it { expect(subject.expense).to eq 8.8 }
    end
  end

  describe '#income' do
    before do
      create(:item, :income, :uah, sum: 1.1)
      create(:item, :income, :usd, sum: 2.2)
      create(:item, :income, :eur, sum: 3.3)

      create(:item, :expense, :uah, sum: 4.4)
    end

    context 'when currency is `uah`' do
      let(:params) { { currency: 'uah' } }

      before { create(:item, :income, :uah, sum: 5.5) }

      it { expect(subject.income).to eq 6.6 }
    end

    context 'when currency is `usd`' do
      let(:params) { { currency: 'usd' } }

      before { create(:item, :income, :usd, sum: 5.5) }

      it { expect(subject.income).to eq 7.7 }
    end

    context 'when currency is `eur`' do
      let(:params) { { currency: 'eur' } }

      before { create(:item, :income, :eur, sum: 5.5) }

      it { expect(subject.income).to eq 8.8 }
    end
  end

  describe '#items' do
    let(:currency) { 'uah' }

    let(:params) { { currency:, year: 2022, month: 10 } }

    before do
      create(:item, date: '2022-01-01', currency: 'uah')
      create(:item, date: '2022-01-01', currency: 'usd')
    end

    context 'when date is `October, 2022` and currency is `uah`' do
      let!(:item_n1) { create(:item, currency: 'uah', date: '2022-10-03') }
      let!(:item_n2) { create(:item, currency: 'uah', date: '2022-10-30') }

      it { expect(subject.items).to eq [item_n2, item_n1] }
    end

    context 'when date is `October, 2022` and currency is `usd`' do
      let(:currency) { 'usd' }

      let!(:item_n3) { create(:item, currency: 'usd', date: '2022-10-03') }
      let!(:item_n4) { create(:item, currency: 'usd', date: '2022-10-30') }

      it { expect(subject.items).to eq [item_n4, item_n3] }
    end

    context 'when date is `September, 2022` and currency is `uah`' do
      let(:params) { { currency: 'uah', year: 2022, month: 9 } }

      let!(:item_n5) { create(:item, currency: 'uah', date: '2022-09-03') }
      let!(:item_n6) { create(:item, currency: 'uah', date: '2022-09-30') }

      it { expect(subject.items).to eq [item_n6, item_n5] }
    end
  end

  describe '#at_end' do
    before do
      allow(subject).to receive(:income).and_return(100.99)

      allow(subject).to receive(:expense).and_return(50.88)
    end

    it { expect(subject.at_end).to eq 50.11 }
  end
end
