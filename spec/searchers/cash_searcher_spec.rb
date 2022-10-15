# frozen_string_literal: true

RSpec.describe CashSearcher do
  subject { described_class.new(:relation, :params) }

  it { is_expected.to be_an ActsAsSearchByCurrency }

  describe '#search_by_supercategory' do
    subject { described_class.search(relation, supercategory:) }

    let(:relation) { Cash.all }

    before do
      create(:cash, :cash, id: 1)
      create(:cash, :bonds, id: 2)
      create(:cash, :deposits, id: 3)
    end

    context 'when #supercategory is `cash`' do
      let(:supercategory) { 'cash' }

      its('ids') { is_expected.to eq [1] }
    end

    context 'when #supercategory is `bonds`' do
      let(:supercategory) { 'bonds' }

      its('ids') { is_expected.to eq [2] }
    end

    context 'when #supercategory is `deposits`' do
      let(:supercategory) { 'deposits' }

      its('ids') { is_expected.to eq [3] }
    end

    context 'when #supercategory is not present' do
      let(:supercategory) { '' }

      it { is_expected.to eq relation }
    end
  end
end
