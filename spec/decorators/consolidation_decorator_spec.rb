require 'rails_helper'

RSpec.describe ConsolidationDecorator do
  let(:consolidation) { stub_model Consolidation }

  let(:date) { Date.new 2019, 5, 1 }

  subject { consolidation.decorate context: { date: date } }

  it { should delegate_method(:year).to(:date) }

  it { should delegate_method(:month).to(:date) }

  its(:date) { should eq date }
end
