# frozen_string_literal: true

RSpec.describe Cash, type: :model do
  it { is_expected.to be_an ActsAsHasFormula }

  it { is_expected.to validate_presence_of :name }

  it { is_expected.to validate_uniqueness_of(:name).case_insensitive.scoped_to(:currency) }

  it { is_expected.to validate_presence_of :formula }

  it { is_expected.to validate_presence_of :currency }

  it { is_expected.to define_enum_for(:currency).with_values(%w[uah usd eur]) }

  it { expect(subject.defined_enums).to include 'supercategory' => { 'cash' => 1, 'bonds' => 2 } }

  it { is_expected.to be_versioned }

  describe '.supercategories' do
    subject { described_class.supercategories }

    it do
      I18n.with_locale(:en) do
        expect(subject).to eq 'Cash' => 'cash', 'Bonds' => 'bonds'
      end
    end

    it do
      I18n.with_locale(:ua) do
        expect(subject).to eq 'Готівка' => 'cash', 'Облігації' => 'bonds'
      end
    end
  end
end
