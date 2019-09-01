# frozen_string_literal: true

RSpec.describe Consolidation do
  it { should be_an Item }

  it { should delegate_method(:name).to(:category) }

  it { should delegate_method(:slug).to(:category) }

  it { should delegate_method(:income?).to(:category) }
end
