# frozen_string_literal: true

RSpec.describe ItemSearcher do
  describe '#month' do
    let(:relation) { double }

    let(:params) { { month: '2021-03' } }

    subject { described_class.new relation, params }

    context do
      let(:month) { Month.today }

      before { subject.instance_variable_set :@month, month }

      its(:month) { should eq month }
    end

    context do
      its(:month) { should eq Month.new(2021, 3) }
    end
  end
end
