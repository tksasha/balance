# frozen_string_literal: true

RSpec.describe Cash do
  it { is_expected.to be_an ActsAsHasFormula }

  it { is_expected.to validate_presence_of :name }

  it { is_expected.to validate_uniqueness_of(:name).case_insensitive.scoped_to(:currency) }

  it { is_expected.to validate_presence_of :formula }

  it { is_expected.to validate_presence_of :currency }

  it { is_expected.to define_enum_for(:currency).with_values(uah: 0, usd: 1, eur: 3) }

  it { expect(subject.defined_enums).to include 'supercategory' => { 'cash' => 1, 'bonds' => 2, 'deposits' => 3 } }

  it { is_expected.to be_versioned }

  describe '.supercategories' do
    subject { described_class.supercategories }

    it do
      I18n.with_locale(:en) do
        expect(subject).to eq 'Cash' => 'cash', 'Bonds' => 'bonds', 'Deposits' => 'deposits'
      end
    end

    it do
      I18n.with_locale(:ua) do
        expect(subject).to eq 'Готівка' => 'cash', 'Облігації' => 'bonds', 'Депозити' => 'deposits'
      end
    end
  end

  describe '.favorite' do
    subject { described_class.favorite.ids }

    before do
      create(:cash, id: 1, favorite: true)
      create(:cash, id: 2, favorite: false)
      create(:cash, id: 3, favorite: true)
    end

    it { is_expected.to eq [1, 3] }
  end
end
