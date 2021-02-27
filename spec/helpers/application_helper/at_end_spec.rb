# frozen_string_literal: true

RSpec.describe ApplicationHelper, type: :helper do
  subject { helper }

  describe '#at_end' do
    context do
      before { subject.instance_variable_set :@at_end, :at_end }

      its(:at_end) { should eq :at_end }
    end

    context do
      let(:params) { double }

      before { allow(subject).to receive(:params).and_return(params) }

      before { allow(CalculateAtEndService).to receive(:calculate).with(params).and_return(21.49) }

      its(:at_end) { should eq 21.49 }
    end
  end
end
