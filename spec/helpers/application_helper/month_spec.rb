# frozen_string_literal: true

RSpec.describe ApplicationHelper, type: :helper do
  subject { helper }

  describe '#month' do
    before { allow(subject).to receive(:params).and_return(params) }

    context do
      let(:month) { Month.today }

      before { subject.instance_variable_set :@month, month }

      its(:month) { should eq month }
    end

    context do
      let(:params) { { month: nil } }

      let(:month) { Month.new(2021, 4) }

      before { travel_to Date.new(2021, 4) }

      after { travel_back }

      its(:month) { should eq month }
    end

    context do
      let(:params) { { month: '2021-03' } }

      let(:month) { Month.new(2021, 3) }

      its(:month) { should eq month }
    end
  end
end
