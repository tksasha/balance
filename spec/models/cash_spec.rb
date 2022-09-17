# frozen_string_literal: true

RSpec.describe Cash, type: :model do
  it { is_expected.to be_an ActsAsHasFormula }

  it { is_expected.to validate_presence_of :name }

  it { is_expected.to validate_uniqueness_of(:name).case_insensitive.scoped_to(:currency) }

  it { is_expected.to validate_presence_of :formula }

  it { is_expected.to validate_presence_of :currency }

  it { is_expected.to define_enum_for(:currency).with_values(%w[uah usd rub eur]) }

  it { is_expected.to be_versioned }
end
