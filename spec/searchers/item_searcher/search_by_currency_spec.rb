# frozen_string_literal: true

RSpec.describe ItemSearcher do
  subject { described_class.new :relation, :params }

  it { should be_an ActsAsSearchByCurrency }
end
