# frozen_string_literal: true

RSpec.describe Consolidation do
  it { is_expected.to be_an Item }

  it { is_expected.to delegate_method(:name).to(:category) }

  it { is_expected.to delegate_method(:income?).to(:category) }

  it { is_expected.to delegate_method(:id).to(:category).with_prefix }
end
