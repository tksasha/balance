# frozen_string_literal: true

RSpec.describe CategorySearcher do
  subject { described_class.new :relation, :params }

  it { should be_an ActsAsSearchByCurrency }
end
