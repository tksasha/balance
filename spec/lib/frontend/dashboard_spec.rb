# frozen_string_literal: true

RSpec.describe Frontend::Dashboard do
  subject { described_class.new(params) }

  let(:params) { { currency: 'uah' } }

  describe '#income' do
    subject { described_class.new(params).send(:income) }

    before do
      create(:item, :income, :uah, sum: 1.1)
      create(:item, :income, :usd, sum: 2.2)
      create(:item, :income, :eur, sum: 3.3)

      create(:item, :expense, :uah, sum: 4.4)
    end

    it { is_expected.to be_a(BigDecimal) }

    context 'when currency is `uah`' do
      let(:params) { { currency: 'uah' } }

      before { create(:item, :income, :uah, sum: 5.5) }

      it { is_expected.to eq 6.6 }
    end

    context 'when currency is `usd`' do
      let(:params) { { currency: 'usd' } }

      before { create(:item, :income, :usd, sum: 5.5) }

      it { is_expected.to eq 7.7 }
    end

    context 'when currency is `eur`' do
      let(:params) { { currency: 'eur' } }

      before { create(:item, :income, :eur, sum: 5.5) }

      it { is_expected.to eq 8.8 }
    end
  end

  describe '#expense' do
    subject { described_class.new(params).send(:expense) }

    before do
      create(:item, :expense, :uah, sum: 1.1)
      create(:item, :expense, :usd, sum: 2.2)
      create(:item, :expense, :eur, sum: 3.3)

      create(:item, :income, :uah, sum: 4.4)
    end

    it { is_expected.to be_a(BigDecimal) }

    context 'when currency is `uah`' do
      let(:params) { { currency: 'uah' } }

      before { create(:item, :expense, :uah, sum: 5.5) }

      it { is_expected.to eq 6.6 }
    end

    context 'when currency is `usd`' do
      let(:params) { { currency: 'usd' } }

      before { create(:item, :expense, :usd, sum: 5.5) }

      it { is_expected.to eq 7.7 }
    end

    context 'when currency is `eur`' do
      let(:params) { { currency: 'eur' } }

      before { create(:item, :expense, :eur, sum: 5.5) }

      it { is_expected.to eq 8.8 }
    end
  end

  describe '#items' do
    subject { described_class.new(params).items }

    before do
      create(:item, date: '2022-01-01', currency: 'uah')
      create(:item, date: '2022-01-01', currency: 'usd')
    end

    context 'when date is `October, 2022` and currency is `uah`' do
      let(:params) { { currency: 'uah', year: 2022, month: 10 } }

      let!(:item_n1) { create(:item, currency: 'uah', date: '2022-10-03') }
      let!(:item_n2) { create(:item, currency: 'uah', date: '2022-10-30') }

      it { is_expected.to eq [item_n2, item_n1] }
    end

    context 'when date is `October, 2022` and currency is `usd`' do
      let(:params) { { currency: 'usd', year: 2022, month: 10 } }

      let!(:item_n3) { create(:item, currency: 'usd', date: '2022-10-03') }
      let!(:item_n4) { create(:item, currency: 'usd', date: '2022-10-30') }

      it { is_expected.to eq [item_n4, item_n3] }
    end

    context 'when date is `September, 2022` and currency is `uah`' do
      let(:params) { { currency: 'uah', year: 2022, month: 9 } }

      let!(:item_n5) { create(:item, currency: 'uah', date: '2022-09-03') }
      let!(:item_n6) { create(:item, currency: 'uah', date: '2022-09-30') }

      it { is_expected.to eq [item_n6, item_n5] }
    end
  end

  describe '#at_end' do
    subject { described_class.new(params).at_end }

    let(:params) { { currency: 'uah' } }

    before do
      create(:item, :uah, :income, sum: 100.99)
      create(:item, :uah, :expense, sum: 50.88)
    end

    it { is_expected.to be_a(BigDecimal) }

    it { is_expected.to eq 50.11 }
  end

  describe '#balance' do
    subject { described_class.new(params).balance }

    let(:params) { { currency: 'uah' } }

    before do
      create(:cash, :uah, sum: 100.99)

      create(:item, :uah, :income, sum: 200.99)
      create(:item, :uah, :expense, sum: 99.99)
    end

    it { is_expected.to be_a(BigDecimal) }

    it { is_expected.to eq(-0.01) }
  end

  describe '#item' do
    subject { described_class.new(params).item }

    context 'when currency is `uah`' do
      let(:params) { { currency: 'uah' } }

      it { is_expected.to be_a_new_record }

      it { is_expected.to be_a(Item) }

      it { expect(subject.currency).to eq 'uah' }
    end
  end
end
