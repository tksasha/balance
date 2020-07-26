# frozen_string_literal: true

RSpec.describe Cash, type: :model do
  it { should be_a ActsAsHasFormula }

  it { should validate_presence_of :name }

  it { should validate_uniqueness_of(:name).case_insensitive.scoped_to(:currency) }

  it { should validate_presence_of :formula }

  it { should validate_presence_of :currency }

  it { should define_enum_for(:currency).with_values(%w[uah usd rub]) }

  it { should act_as_paranoid }

  it { should be_versioned }
end
