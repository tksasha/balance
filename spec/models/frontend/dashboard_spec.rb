# frozen_string_literal: true

RSpec.describe Frontend::Dashboard, type: :model do
  subject { described_class.new(currency:, year:, month:) }

  let(:currency) { 'uah' }

  let(:year) { 2022 }

  let(:month) { 10 }

  its(:date) { is_expected.to eq Date.new(2022, 10, 1)..Date.new(2022, 10, 31) }

  describe '#items' do
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
      let(:month) { 9 }

      let!(:item_n5) { create(:item, currency: 'uah', date: '2022-09-03') }
      let!(:item_n6) { create(:item, currency: 'uah', date: '2022-09-30') }

      it { expect(subject.items).to eq [item_n6, item_n5] }
    end
  end
end
