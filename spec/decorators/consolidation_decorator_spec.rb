# frozen_string_literal: true

RSpec.describe ConsolidationDecorator do
  let(:consolidation) { stub_model Consolidation }

  let(:params) { {} }

  subject { consolidation.decorate context: params }

  it { should delegate_method(:year).to(:date) }

  it { should delegate_method(:month).to(:date) }

  describe '#date' do
    let(:params) { { month: '2021-03' } }

    its(:date) { should eq Month.new(2021, 3) }
  end
end
