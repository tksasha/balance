# frozen_string_literal: true

RSpec.describe ExchangeRate, type: :model do
  it { should validate_presence_of(:date) }

  it { should validate_uniqueness_of(:date).scoped_to(%i[from to]) }

  it { should validate_presence_of(:from) }

  it { should validate_presence_of(:to) }

  it { should validate_presence_of(:rate) }

  it { should validate_numericality_of(:rate).is_greater_than(0) }

  it { should define_enum_for(:from).with_values(%w[uah usd rub]).with_suffix }

  it { should define_enum_for(:to).with_values(%w[uah usd rub]).with_suffix }

  context do
    subject { described_class }

    its(:default_per_page) { should eq 10 }
  end
end
