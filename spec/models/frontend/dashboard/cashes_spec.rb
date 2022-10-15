# frozen_string_literal: true

RSpec.describe Frontend::Dashboard::Cashes do
  subject { described_class.new(currency:) }

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

  describe '#all' do
    before do
      create(:cash, :cash, :uah, id: 1, name: 'A UAH')
      create(:cash, :cash, :usd, id: 2, name: 'A USD')
      create(:cash, :cash, :eur, id: 3, name: 'A EUR')

      create(:cash, :bonds, :uah, id: 4, name: 'B UAH')
      create(:cash, :bonds, :usd, id: 5, name: 'B USD')
      create(:cash, :bonds, :eur, id: 6, name: 'B EUR')

      create(:cash, :deposits, :uah, id: 7, name: 'C UAH')
      create(:cash, :deposits, :usd, id: 8, name: 'C USD')
      create(:cash, :deposits, :eur, id: 9, name: 'C EUR')
    end

    context 'when currency is `uah`' do
      let(:currency) { 'uah' }

      its('all.ids') { is_expected.to eq [1, 4, 7] }
    end

    context 'when currency is `usd`' do
      let(:currency) { 'usd' }

      its('all.ids') { is_expected.to eq [2, 5, 8] }
    end

    context 'when currency is `eur`' do
      let(:currency) { 'eur' }

      its('all.ids') { is_expected.to eq [3, 6, 9] }
    end
  end

  describe '.all' do
    let(:instance) { double }

    before do
      allow(described_class).to receive(:new).and_return(instance)

      allow(instance).to receive(:all)

      described_class.all(currency: 'uah')
    end

    it { expect(described_class).to have_received(:new).with(currency: 'uah') }

    it { expect(instance).to have_received(:all) }
  end

  describe '#summary' do
    before do
      create(:cash, :uah, :cash, sum: 1.12)
      create(:cash, :uah, :bonds, sum: 1.34)
      create(:cash, :uah, :deposits, sum: 1.56)

      create(:cash, :usd, :cash, sum: 2.12)
      create(:cash, :usd, :bonds, sum: 2.34)
      create(:cash, :usd, :deposits, sum: 2.56)

      create(:cash, :eur, :cash, sum: 3.12)
      create(:cash, :eur, :bonds, sum: 3.34)
      create(:cash, :eur, :deposits, sum: 3.56)
    end

    context 'when currency is `uah`' do
      let(:currency) { :uah }

      its(:summary) { is_expected.to eq 'cash' => 1.12, 'bonds' => 1.34, 'deposits' => 1.56 }
    end

    context 'when currency is `usd`' do
      let(:currency) { :usd }

      its(:summary) { is_expected.to eq 'cash' => 2.12, 'bonds' => 2.34, 'deposits' => 2.56 }
    end

    context 'when currency is `eur`' do
      let(:currency) { :eur }

      its(:summary) { is_expected.to eq 'cash' => 3.12, 'bonds' => 3.34, 'deposits' => 3.56 }
    end
  end
end
