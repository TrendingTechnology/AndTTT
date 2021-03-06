package com.github.dawidd6.andttt.animations;

import android.view.View;
import android.view.animation.Transformation;

public class LightenAnimation extends BaseAnimation {
    public LightenAnimation(View view, int duration) {
        super(view, duration);
        start();
    }

    @Override
    protected void applyTransformation(float interpolatedTime, Transformation t) {
        float alpha = 0.8f - (0.8f * interpolatedTime);
        view.setAlpha(alpha);
    }
}
