# frozen_string_literal: true

RSpec.describe CashSearcher do
  subject { described_class.new :relation, :params }

  it { is_expected.to be_an ActsAsSearchByCurrency }
end
