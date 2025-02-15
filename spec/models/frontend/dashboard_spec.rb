# frozen_string_literal: true

RSpec.describe Frontend::Dashboard do
  subject { described_class.new(params) }

  let(:currency) { 'usd' }
  let(:month) { '2023-05' }
  let(:params) { { currency:, month: } }

  describe '#currency', skip: 'private method' do
    its(:currency) { is_expected.to eq 'usd' }

    context 'when currency is nil' do
      let(:params) { {} }

      its(:currency) { is_expected.to eq 'uah' }
    end
  end

  describe '#month', skip: 'private method' do
    its(:month) { is_expected.to eq Month.new(2023, 5) }

    context 'when month is nil' do
      let(:params) { {} }

      before { travel_to '2023-12-31' }

      its(:month) { is_expected.to eq Month.new(2023, 12) }
    end
  end

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

  describe '#consolidations' do
    let(:currency) { 'eur' }
    let(:month) { Month.new(2023, 5) }
    let(:params) { { currency:, month: '2023-05' } }

    before do
      allow(Frontend::Reports::Consolidations).to receive(:call)

      subject.consolidations
    end

    it { expect(Frontend::Reports::Consolidations).to have_received(:call).with(currency:, month:) }
  end

  describe '#cashes' do
    let(:currency) { 'eur' }

    before do
      allow(Frontend::Reports::Cashes).to receive(:call)

      subject.cashes
    end

    it { expect(Frontend::Reports::Cashes).to have_received(:call).with(currency:) }
  end

  describe '#cashes_sum', skip: 'private method' do
    let(:currency) { 'eur' }

    before do
      create(:cash, currency: 'uah', sum: 11.11)
      create(:cash, currency: 'usd', sum: 22.22)
      create(:cash, currency: 'eur', sum: 33.33)
      create(:cash, currency: 'eur', sum: 44.44)
    end

    its(:cashes_sum) { is_expected.to eq 77.77 }
  end
end
