# frozen_string_literal: true

RSpec.describe ItemSearcher do
  subject { described_class.search(Item.all, params) }

  describe '#search_by_currency' do
    subject { described_class.new(double, {}) }

    it { is_expected.to be_an(ActsAsSearchByCurrency) }
  end

  describe '#search_by_month' do
    before do
      travel_to '2023-05-01'

      create(:item, id: 1, date: '2023-01-01')
      create(:item, id: 2, date: '2023-05-01')
      create(:item, id: 3, date: '2023-05-01')
      create(:item, id: 4, date: '2023-01-01')
    end

    context 'when month was not specified' do
      let(:params) { {} }

      its(:ids) { is_expected.to match_array [2, 3] }
    end

    context 'when month was specified' do
      let(:params) { { month: '2023-01' } }

      its(:ids) { is_expected.to match_array [1, 4] }
    end
  end

  describe '#search_by_category' do
    before do
      create(:category, id: 1)
      create(:category, id: 2)

      create(:item, id: 1, category_id: 1)
      create(:item, id: 2, category_id: 1)
      create(:item, id: 3, category_id: 2)
      create(:item, id: 4, category_id: 2)
    end

    context 'when category_id was not specified' do
      let(:params) { {} }

      its(:ids) { is_expected.to match_array [1, 2, 3, 4] }
    end

    context 'when category_id is blank' do
      let(:params) { { category_id: '' } }

      its(:ids) { is_expected.to match_array [1, 2, 3, 4] }
    end

    context 'when category_id is specified' do
      let(:params) { { category_id: 2 } }

      its(:ids) { is_expected.to match_array [3, 4] }
    end
  end
end
